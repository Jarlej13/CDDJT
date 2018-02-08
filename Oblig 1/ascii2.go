package main

import "fmt"
const ascii = "\x80\xF7\xBE"



func main() {
	IterateOverASCIIStringLiteral()
	ExtendedASCIIText(ascii)
}

func IterateOverASCIIStringLiteral() {
	a := ""
	b := "\\0x"
	for i := 128; i < 256; i++ {
		h := fmt.Sprintf("%X", i)
		st := string(i)
		a = a + b + h
		fmt.Printf("%s %s %b\n", h, st, i)
	}
}

func ExtendedASCIIText(tekst string) {
	a := tekst
	var b  = fmt.Sprintf("%c", a)
	for i := 0; i < len(a); i++ {
		fmt.Printf("%c", a[i])
	}
	return b
}