package adder

import (
	"io/ioutil"
	"fmt"
	"strings"
	"strconv"
	"log"
)

func Adder(done chan bool) {
	b, err := ioutil.ReadFile("tempfile.txt.lock")
	//Her kommer en feilmelding dersom filen den skal lese ikke finnes.
	//Feilmeldingen sier "The system cannot find the file specified"
	if err != nil {
		fmt.Print(err)
	}

	inputS := string(b)

	tallS := strings.Fields(inputS)

	tall1S := string(tallS[0])
	tall2S := string(tallS[1])

	tall1, err := strconv.Atoi(tall1S)
	//Her kommer en feilmelding dersom tall1S ikke er et tall.
	//Feilmeldingen sier "Invalid syntax"
	if err != nil {
		log.Fatal(err)
	}

	tall2, err := strconv.Atoi(tall2S)
	//Her kommer en feilmelding dersom tall2S ikke er et tall.
	//Feilmeldingen sier "Invalid syntax"
	if err != nil {
		log.Fatal(err)
	}

	sum := tall1 + tall2

	sumS := strconv.Itoa(sum)

	sumB := []byte(sumS)

	ioutil.WriteFile("tempfile.txt.lock", sumB, 0777)

	done <- true
}
