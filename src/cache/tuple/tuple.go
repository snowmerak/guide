package tuple

type Tuple[A any, B any] struct {
	a A
	b B
}

func New[A any, B any](a A, b B) Tuple[A, B] {
	return Tuple[A, B]{a, b}
}

func (t Tuple[A, B]) A() A {
	return t.a
}

func (t Tuple[A, B]) B() B {
	return t.b
}

func (t Tuple[A, B]) Unpack() (A, B) {
	return t.a, t.b
}
