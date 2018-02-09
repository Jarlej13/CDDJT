package main	

import (
"fmt"
"os"
"os/signal"

)
// Prossessen bruker 50% av prosessor og 53% av minne på en datamaskin.  Avslutningsmeldinger som programmet håndterer og godkjenner er følgende: Ctrl-C og Ctrl-break
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
