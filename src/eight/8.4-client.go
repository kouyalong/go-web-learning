package main

import (
	"fmt"
	"log"
	"net/rpc"
	"os"
)

type Args struct {
	A, B int
}


type Quotient struct {
	Quo, Rem int
}

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Usage: ", os.Args[0], "server")
	}

	serverAddress := os.Args[1]
	client, err := rpc.DialHTTP("tcp", serverAddress + ":50051")
	if err != nil {
		log.Fatal("dialing", err)
	}

	args := Args{17, 8}

	var reply int
	err = client.Call("Arith.Multiply", args, &reply)
	if err != nil {
		log.Fatal("arith error: ", err)
	}

	fmt.Printf("Arith: %d * %d=%d\n", args.A, args.B, reply)

	var quto Quotient
	err = client.Call("Arith.Divide", args, &quto)
	if err != nil {
		log.Fatal("arith error: ", err)
	}
	fmt.Printf("Arith: %d / %d = %d  remainder %d \n", args.A, args.B, quto.Quo, quto.Rem)
}