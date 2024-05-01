package hashing

import (
	"fmt"
	"testing"
)

func Test_MyOpenAddressing(t *testing.T) {

	openAdr := new(MyHashOpenAddressing).Init(7)

	openAdr.Insert(0)
	openAdr.Insert(1)
	openAdr.Insert(2)
	openAdr.Insert(3)
	openAdr.Insert(4)
	openAdr.Insert(5)
	openAdr.Insert(6)
	openAdr.Insert(7)
	openAdr.Insert(8)

	i := openAdr.Search(28)

	fmt.Print(i)
}
