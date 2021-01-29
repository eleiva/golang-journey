package examples

type Array struct {}

func (t *Array) Foo() string {
	return "foo1"
}

func (t *Array) Second() string {
	return "second"
}

func (t *Array) InitializationAnIntegerArray() []int {
	setA := []int{1, 2, 3,4,5}
	return setA
}
