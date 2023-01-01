package main

import (
	"fmt"
	"go/ast"
	"go/format"
	"go/parser"
	"go/token"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"sort"
	"strings"
)

var HEADER = `
// Code generated by github.com/barrettj12/collections/internal/gen. DO NOT EDIT.
// Generated from package "strings", version %s
package collections

import "strings"
`[1:]

// Generate String methods from Go's strings library.
func main() {
	// Open file to be written
	outFile, err := os.Create("string_gen.go")
	if err != nil {
		log.Fatalf("couldn't create output file: %v", err)
	}
	defer outFile.Close()

	outFile.WriteString(fmt.Sprintf(HEADER, goVersion()))

	// Parse strings package
	fset := token.NewFileSet()
	pkgPath := filepath.Join(runtime.GOROOT(), "src/strings")
	pkgs, err := parser.ParseDir(fset, pkgPath, nil, parser.ParseComments)
	if err != nil {
		log.Fatalf("couldn't parse strings package: %v", err)
	}

	// Sort files to get deterministic output
	files := make([]string, 0, len(pkgs["strings"].Files))
	for filename := range pkgs["strings"].Files {
		// ignore test files
		if strings.HasSuffix(filename, "_test.go") {
			continue
		}
		files = append(files, filename)
	}
	sort.Strings(files)

	for _, filename := range files {
		file := pkgs["strings"].Files[filename]

		// remove non-exported declarations
		ast.FileExports(file)

		for _, decl := range file.Decls {
			// only match function declarations
			funcDecl, ok := decl.(*ast.FuncDecl)
			if !ok {
				continue
			}

			// We want to make two different transformations to strings functions.
			//
			// Declarations of the form
			//    func Foo(s string, a A, b B, ...) string
			// should be transformed to
			//    func (s *String) Foo(a A, b B, ...)
			//
			// while declarations of the form
			//    func Foo(s string, a A, b B, ...) R
			// where R != string, should be transformed to
			//    func (s *String) Foo(a A, b B, ...) R

			// check there is no receiver (this is a function)
			if funcDecl.Recv != nil {
				continue
			}

			// check first arg exists and has type string
			args := funcDecl.Type.Params
			if len(args.List) == 0 {
				continue
			}
			firstArgType, ok := args.List[0].Type.(*ast.Ident)
			if !ok || firstArgType.Name != "string" {
				continue
			}

			// The function matches, so transform it into a method, as described
			// above.

			// Add receiver to function
			funcDecl.Recv = &ast.FieldList{
				List: []*ast.Field{{
					Names: []*ast.Ident{{Name: "s"}},
					Type:  &ast.Ident{Name: "*String"},
				}},
			}

			// Remove first arg
			if len(args.List[0].Names) == 1 {
				args.List = args.List[1:]
			} else {
				args.List[0].Names = args.List[0].Names[1:]
			}

			// If this function returns a single string, remove this return value,
			// and make it a "mutating method" on String.
			// Otherwise, keep the return values the same, and make this a
			// "non-mutating" method on String.
			mutating := false
			// Certain methods don't make sense as mutating, even though they match
			// the type.
			funcName := funcDecl.Name.Name
			if funcName != "Repeat" && funcName != "Clone" {
				res := funcDecl.Type.Results
				if res != nil && res.NumFields() == 1 {
					switch retType := res.List[0].Type.(type) {
					case *ast.Ident:
						if retType.Name == "string" {
							mutating = true
							funcDecl.Type.Results = nil
						}
					case *ast.StarExpr:
						baseType, ok := retType.X.(*ast.Ident)
						if ok && baseType.Name == "Reader" {
							// need to qualify package name
							baseType.Name = "strings.Reader"
						}
					}
				}
			}

			// Delete body - we will write this later
			funcDecl.Body = nil

			// For mutating functions, make some heuristic changes to the comment.
			if mutating {
				alterComment(funcDecl.Doc)
			}

			// Write function signature to file
			outFile.WriteString("\n")
			format.Node(outFile, fset, funcDecl)

			// Write function body
			outFile.WriteString(" {\n")
			if mutating {
				outFile.WriteString(fmt.Sprintf("	*s = String(strings.%s(string(*s)%s))\n",
					funcName, argsString(args)))
			} else {
				outFile.WriteString(fmt.Sprintf("	return strings.%s(string(*s)%s)\n",
					funcName, argsString(args)))
			}
			outFile.WriteString("}\n")
		}
	}
}

func argsString(args *ast.FieldList) string {
	strArgs := make([]string, 1, args.NumFields()+1)
	for _, field := range args.List {
		for _, name := range field.Names {
			strArgs = append(strArgs, name.Name)
		}
	}
	return strings.Join(strArgs, ", ")
}

// Replace some common phrases in comments, so that the resulting comment
// more accurately describes the mutating methods.
func alterComment(doc *ast.CommentGroup) {
	if doc != nil && len(doc.List) > 0 {
		doc.List[0].Text = strings.Replace(doc.List[0].Text,
			"returns s with ", "modifies s so that it has ", 1)
		doc.List[0].Text = strings.Replace(doc.List[0].Text,
			"returns a copy of the string s with ", "modifies s so that it has ", 1)
		doc.List[0].Text = strings.Replace(doc.List[0].Text,
			"returns a slice of the string s with ", "modifies s so that it has ", 1)
		doc.List[0].Text = strings.Replace(doc.List[0].Text,
			"returns a slice of the string s, with ", "modifies s so that it has ", 1)
		doc.List[0].Text = strings.Replace(doc.List[0].Text,
			"returns s without ", "removes from s ", 1)
	}
}

func goVersion() string {
	b, err := exec.Command("go", "version").Output()
	if err != nil {
		return "(unknown)"
	}
	return strings.Fields(string(b))[2]
}