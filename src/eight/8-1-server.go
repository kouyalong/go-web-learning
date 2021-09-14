package main

import (
	"fmt"
	"net"
	"os"
	"strconv"
	"time"
)

func checkErr(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error: %s", err.Error())
	}
}


func handlerClient(conn net.Conn) {
	//conn.SetReadDeadline(time.Now().Add(10 * time.Second))
	request := make([]byte, 128)
	defer conn.Close()

	for {
		readLen, err := conn.Read(request)
		if err != nil {
			fmt.Println(err)
			break
		}

		if readLen == 0 {
			break
		}
		str := string(request[0:readLen])

		flag := str == "quit"

		fmt.Println("Client 发送的信息：", str, flag)

		if str == "timestamp" {
			daytime := strconv.FormatInt(time.Now().Unix(), 10)
			conn.Write([]byte(daytime))
		} else if flag == true {
			conn.Write([]byte("quit"))
			fmt.Println("flag is:", flag)
			break
		} else {
			daytime := time.Now().String()
			fmt.Println("str is:", str)
			conn.Write([]byte(daytime))
		}

		request = make([]byte, 128)
	}
	fmt.Printf("客户端断开连接 %s", conn.RemoteAddr())
}

func main() {
	service := ":8888"
	tcpAddr, err := net.ResolveTCPAddr("tcp4", service)
	checkErr(err)

	listener, err := net.ListenTCP("tcp", tcpAddr)
	checkErr(err)

	for {
		conn, err := listener.Accept()
		if err != nil {
			continue
		}

		go handlerClient(conn)
	}
}