package main

import (
"fmt"
"os"
"os/signal"

)
// Prossessen bruker ca 50% av prosessor og 12% av minne på en datamaskin. 30% prosessor og 16% på en annen.
// Minne er uforandret om man kjører programmet eller ikke. 
//Avslutningsmeldinger som programmet håndterer og godkjenner er følgende: Ctrl-C og Ctrl-break
func main() {
	d := make(chan os.Signal, 2)
	signal.Notify(d, os.Interrupt,)
	go func() {
		<-d
		fmt.Println("EN AVSLUTSNINGSMELDING VEL")
		os.Exit(1)
	}()


	for i := 0;;i++{
		fmt.Println(i)
	}
}
