package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

func checkError(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error: %s", err.Error())
	}
}

func main() {
	if len(os.Args) != 2 {
		fmt.Fprintf(os.Stderr, "Usage: %s host:port", os.Args[0])
		os.Exit(1)
	}

	service := os.Args[1]
	fmt.Println(service)
	tcpAdder, err := net.ResolveTCPAddr("tcp4", service)
	checkError(err)

	conn, err := net.DialTCP("tcp", nil, tcpAdder)
	checkError(err)
	defer conn.Close()

	fmt.Println("Start message with Server!")

	buff := make([]byte, 256)

	reader := bufio.NewReader(os.Stdin)

	for {
		lineBytes, _, _ := reader.ReadLine()
		conn.Write(lineBytes)

		n, err := conn.Read(buff)
		if err != nil {
			checkError(err)
			break
		}
		serverMsg := string(buff[0:n])
		fmt.Println("Server 的信息：", serverMsg)
		if serverMsg == "quit" {
			break
		}
	}
	os.Exit(0)
}