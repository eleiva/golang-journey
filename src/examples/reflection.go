package examples

import "reflect"

type Reflection struct {}

func (r * Reflection) GetStringTypeUsingReflection()  interface{} {
	value := "something" // Change this assignment to get different results
	return reflect.TypeOf(value)
}

func (r * Reflection) GetArrayTypeUsingReflection()  interface{} {
	value := []int{1,4} // Change this assignment to get different results
	return reflect.TypeOf(value)
}

func (r * Reflection) GetFloatTypeUsingReflection()  interface{} {
	value := 22.3 // Change this assignment to get different results
	return reflect.TypeOf(value)
}