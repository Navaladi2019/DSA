package main

import (
	"fmt"
	"reflect"
	"strings"
)

func main() {

	fmt.Println(isIsomorphic("bbbaaaba", "aaabbbba"))
	x := 10
	v := reflect.ValueOf(x)

	fmt.Println(v)
	k := v.Kind()
	fmt.Println(k == reflect.Int) // outputs "int"

}

func isIsomorphic(s string, t string) bool {

	if len(s) != len(t) {
		return false
	}

	for i, _ := range s {

		if strings.Index(s, string(s[i])) != strings.Index(t, string(t[i])) {
			return false
		}

	}

	return true
}
