package main

import (
	"os"
	"io/ioutil"
	"fmt"
	"os/signal"
	"awesomeProject1/Oblig2/adder"
)

func main() {
	d := make(chan os.Signal, 2)
	signal.Notify(d, os.Interrupt,)
	go func() {
		<-d
		fmt.Println("Prosess avsluttet")
		os.Exit(1)
	}()

	input1 := os.Args[1]
	input2 := os.Args[2]
	input := input1 + " " + input2

	numbers(input)
}

func numbers(input string) {
	done := make(chan bool, 1)

	inputB := []byte(input)

	ioutil.WriteFile("tempfile.txt.lock", inputB, 0777)

	go adder.Adder(done)

	<-done

	b, err := ioutil.ReadFile("tempfile.txt.lock")
	//Her kommer en feilmelding dersom filen den skal lese ikke finnes.
	//Feilmeldingen sier "The system cannot find the file specified"
	if err != nil {
		fmt.Print(err)
	}

	resultat := string(b)

	fmt.Println(resultat)
}
