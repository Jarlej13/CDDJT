package main

import (
	"os"
	"strconv"
	"fmt"
	"log"
	"os/signal"
)

func main() {
	d := make(chan os.Signal, 2)
	signal.Notify(d, os.Interrupt,)
	go func() {
		<-d
		fmt.Println("Prosess avsluttet")
		os.Exit(1)
	}()

	functionA()
}

func functionA() {
	input1 := os.Args[1]
	input2 := os.Args[2]

	tall1, err := strconv.Atoi(input1)
	//Her kommer en feilmelding dersom det første argumentet ikke er et tall.
	//Feilmeldingen sier "Invalid syntax"
	if err != nil {
		log.Fatal(err)
	}

	tall2, err := strconv.Atoi(input2)
	//Her kommer en feilmelding dersom det andre argumentet ikke er et tall.
	//Feilmeldingen sier "Invalid syntax"
	if err != nil {
		log.Fatal(err)
	}

	c := make(chan int)

	go functionB(tall1, tall2, c)

	resultat := <-c

	fmt.Println(resultat)
}

func functionB(tall1 int, tall2 int, c chan int) {
	resultat := tall1 + tall2

	c <- resultat
}
