package main

import (
	"fmt"
	"reflect"
)

func main() {
	x := 10
	v := reflect.ValueOf(x)

	fmt.Println(v)
	k := v.Kind()
	fmt.Println(k == reflect.Int) // outputs "int"

}
