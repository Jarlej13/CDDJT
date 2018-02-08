package main

import (
	"testing"
	"fmt"
	"strconv"
)

func TestExtendedASCIIText(t *testing.T) {

	a := fmt.Sprintf("%d", ascii)
	i, err := strconv.Atoi(a)

	if i <  {
		fmt.Println("ERROR, value not in extended ASCII")
	}
	/*if !reflect.DeepEqual(values, expected) {
		t.Fatalf("expected %d, actual is %d", 1, values[0])
	}*/
}