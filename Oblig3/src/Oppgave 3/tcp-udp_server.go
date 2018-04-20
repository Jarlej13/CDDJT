package main

import (
	"net"
	"fmt"
	"bufio"
	"os"
	"time"
)


func CheckError(err error) {
	if err  != nil {
		fmt.Println("Error: " , err)
		os.Exit(0)
	}
}
var quotes = []string{
	`Education is the most powerful weapon which you can change the world. - Nelson Mandela
	`,
}


func tcp2(conn net.Conn) {
	fmt.Println("Handling new connection...")



	//close connection when this function ends
	defer func() {
		fmt.Println("Closing connection...")
		conn.Close()
	}()

	timeoutDuration := 5 * time.Second
	bufReader := bufio.NewReader(conn)
	// run loop forever (or until ctrl-c)
	for {
		conn.SetReadDeadline(time.Now().Add(timeoutDuration))
		message, err := bufReader.ReadString('\n')
		if err != nil {
			fmt.Println(err)
			return
		}
		newMessage := string(message)
		fmt.Println(newMessage)
		conn.Write([]byte(quotes[0]))
	}
}

func tcp() {
	listener, err := net.Listen("tcp", ":17")
	if err != nil {

		fmt.Printf("Some error %v", err)
		return
	}

	defer func() {
		listener.Close()
		fmt.Println("listener closed")
	}()

	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println(err)
			break
		}

		go tcp2(conn)
	}
}

func main () {
	//time.Sleep(1*time.Second)
	go tcp()
	ServerAddr, err := net.ResolveUDPAddr("udp", ":17")
	CheckError(err)

	/* Now listen at selected port */
	ServerConn, err := net.ListenUDP("udp", ServerAddr)
	CheckError(err)
	defer ServerConn.Close()

	buf := make([]byte, 512)

	for {
		n, addr, err := ServerConn.ReadFromUDP(buf)
		fmt.Println("Received ", string(buf[0:n]))
		//go quotes()
		// send new string back to client
		ServerConn.WriteToUDP([]byte(quotes[0]), addr)

		if err != nil {
			fmt.Println("Error: ", err)
		}
	}
}