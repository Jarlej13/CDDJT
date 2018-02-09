package main

import (
	"testing"
	"fmt"
	//"strconv"
)

func TestExtendedASCIIText(t *testing.T) {
	for i := 0; i<len(ascii);i++ {
		if (ascii[i] < 128) {
			t.Fail()
			fmt.Println("ERROR, value not in extended ASCII")
		}
	}
}