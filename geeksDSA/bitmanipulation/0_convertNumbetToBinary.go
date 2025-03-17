package bitmanipulation

import (
	"fmt"
	"math"
)

func ConvertToBinary(num int) string {

	bits := ""

	if num == 0 {
		bits = "0"
	}

	for num > 0 {

		if num%2 == 1 {
			bits += "1"
		} else {
			bits += "0"
		}

		num = num / 2
	}

	bits = reverseBits(bits)
	return bits
}

func ConvertBinaryToNumber(bits string) float64 {
	var res float64 = 0
	length := len(bits)

	fmt.Println(length)

	for i := len(bits) - 1; i >= 0; i-- {

		curr := 0

		if bits[i] == '1' {
			curr = 1
		}

		pow := (len(bits) - 1 - i)
		res += (float64(curr) * (math.Pow(float64(2), float64(pow))))
	}

	fmt.Println(res)

	return res
}

func reverseBits(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

func oneAndtwoComplement(n int) {

	biist := ConvertToBinary(n)

	r := []rune(biist)

	for i := 0; i < len(r); i++ {

		if r[i] == '1' {
			r[i] = '0'
		} else {
			r[i] = '1'
		}
	}

	r[len(r)-1] = '1'
	str := string(r)
	ConvertBinaryToNumber(str)
}
