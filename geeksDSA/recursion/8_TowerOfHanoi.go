package recursion

import "fmt"

func TOH(n int, A rune, B rune, C rune) {

	if n == 1 {

		fmt.Printf("Moving %d from %c to %c \n", n, A, C)
		return
	}
	fmt.Println(string(A), string(B), string(C), n)
	TOH(n-1, A, C, B)

	fmt.Printf("Moving %d from %c to %c \n", n, A, C)
	fmt.Println(string(A), string(B), string(C), n)
	TOH(n-1, B, A, C)
}
