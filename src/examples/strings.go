package examples

import (
	"math/rand"
	"strconv"
)

type String struct {}

func (s * String) ConvertAnIntToString() string {
	number := rand.Intn(100)
	i := strconv.Itoa(number)

	return i
}

func  (s * String) ConvertStringToInt() int {
	str := "2"
	i, _ := strconv.Atoi(str)

	return i
}
