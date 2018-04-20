package main

import "net"
import "fmt"
import "bufio"
import "os"
import (
	"sync"
	"time"
)


func wait(seconds int, wg* sync.WaitGroup) {
	defer wg.Done()

	time.Sleep(time.Duration(seconds) * time.Second)
	fmt.Println("Slept ", seconds, " seconds ..")
}

func main() {
	// connect to this socket
	conn, _ := net.Dial("tcp", "localhost:17")
	// read in input from stdin
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Print("What's your message of utter doom? ")
		text, _ := reader.ReadString('\n')
		// send to socket
		fmt.Fprintf(conn, text + "\n")
		// listen for reply
		message, _ := bufio.NewReader(conn).ReadString('\n')
		fmt.Print("Message from server: "+message)
	}
}