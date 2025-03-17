package bitmanipulation

import (
	"fmt"
	"testing"
)

func Test_ConvertToBinary(t *testing.T) {
	bin := ConvertToBinary(13)
	fmt.Println(ConvertBinaryToNumber(bin))
	oneAndtwoComplement(13)
}
