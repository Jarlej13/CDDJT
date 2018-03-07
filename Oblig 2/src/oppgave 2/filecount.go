package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"bufio"
	"strings"
)

var antall []int
var biggest, n, m, k int
func main() {
	filnavn := os.Args[1]


	b, err := ioutil.ReadFile(filnavn) // just pass the file name
	if err != nil {
		fmt.Print(err)
	}

	//fmt.Println(b) // print the content as 'bytes'

	str := string(b) // convert content to a 'string'

	//fmt.Println(str) // print the content as a 'string'

	for i:= 32;  i <= 126; i++{
		iterate(i, str);
	}
	fmt.Println("Information about", filnavn, ":")
	fmt.Println("")
	lineCounter(filnavn);
	fmt.Println("")
	fmt.Println("Most common runes:")
	fmt.Println("")

	for j := 0; j < 5; j++{
		biggest = 0
		for l := 0; l < len(antall); l++ {
			n = antall[l]
			if n > biggest {
				biggest = n
				m = l + 32
				k = l
			}
		}
		antall[k] = 0
		st := string(m)
		st1 := fmt.Sprintf("%s", st)
		fmt.Println(j+1, ".", "Rune:", st1, ",", "Counts:",   biggest)
	}
}

func iterate(tall int, tekst string){
	st := string(tall)
	symbol := fmt.Sprintf("%s", st)
	antallet := strings.Count(tekst, symbol)
	antall = append(antall, antallet)

}

func lineCounter(filnavn string) {
	file, _ := os.Open(filnavn)
	fileScanner := bufio.NewScanner(file)
	lineCount := 0
	for fileScanner.Scan() {
		lineCount++
	}
	fmt.Println("number of lines:", lineCount)
}
