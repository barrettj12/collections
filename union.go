package collections

type OneOf[T, U any] struct {
	t *T
	u *U
}

func Inject0[T, U any](t T) OneOf[T, U] {
	return OneOf[T, U]{
		t: &t,
	}
}

func Inject1[T, U any](u U) OneOf[T, U] {
	return OneOf[T, U]{
		u: &u,
	}
}

func (o *OneOf[T, U]) Is0() bool {
	return o.t != nil
}

func (o *OneOf[T, U]) Is1() bool {
	return o.u != nil
}

func (o *OneOf[T, U]) Get0() T {
	return *o.t
}

func (o *OneOf[T, U]) Get1() U {
	return *o.u
}
