package main

import "fmt"
const ascii = "\x80\xF7\xBE"
const ascii2 = "\x22\x20\x80\x20\xF7\x20\xBE\x20\x64\x6F\x6C\x6C\x61\x72\x20\x22"


/*
I denne koden får vi ikke printet ut verdiene x80-x9F. Dette er fordi våre pcer ikke har riktig symbolsett for å skrive disse.
Dette kan fikses med å implementere runer i koden, men vi mener dette er et unødvendig steg i dette eksperimentet.
Vi har også testet på en MAC, hvor det printer som det skal.
Alle medlemmer har testet kode på sin PC og får identisk utfall. Eneste unntak er MAC pc som en på gruppa har hjemme

Vi implementerte 2 metoder for å printe extendedasciitext metoden i oppgave 4B.


 */


func main() {
	IterateOverASCIIStringLiteral()
	ExtendedASCIIText(ascii)
	ExtendedASCIIText2(ascii2)
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
	//fmt.Println(a)
}

var extAscii string

func ExtendedASCIIText(tekst string) (string) {
	extAscii = ""
	a := tekst
	fmt.Sprintf("%c", a)
	fmt.Printf("%s %c %c %c %s %s", "\"" , a[0], a[1], a[2], "dollar","\"")

	for i := 0; i<len(ascii);i++ {
		if (ascii[i] >= 128) {
			extAscii = extAscii + fmt.Sprintf("%c", ascii[i])
		}
	}
	fmt.Println("De extended verdiene er: " + extAscii)
	return extAscii

}

func ExtendedASCIIText2(tekst string) (string){
	extAscii = ""
	a := tekst
	for i := 0; i < len(a); i++ {
		fmt.Printf("%c", a[i])
	}

	for i := 0; i < len(ascii); i++ {
		if (ascii[i] >= 128) {
			extAscii = extAscii + fmt.Sprintf("%c", ascii[i])
		}
	}
	fmt.Println("De extended verdiene er: " + extAscii)
	return extAscii
}
