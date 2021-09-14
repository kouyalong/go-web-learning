package main

import (
	"fmt"
	"os"
)

func dirOp() {
	os.Mkdir("kyl", 0777)
	os.MkdirAll("kyl/t1/t2", 0777)
	err := os.Remove("kyl")
	if err != nil {
		fmt.Println(err)
	}

	os.RemoveAll("kyl")
}

func fileOp() {
	userFile := "kyl.txt"
	fout, err := os.Create(userFile)
	if err != nil {
		fmt.Println(userFile, err)
	}
	defer fout.Close()

	for i := 0; i < 10; i++ {
		fout.WriteString("Just a test A!\r\n")
		fout.Write([]byte("Just a test B!\r\n"))
	}

	fl, err := os.Open(userFile)
	if err != nil {
		fmt.Println(userFile, err)
		return
	}
	defer fl.Close()

	buf := make([]byte, 1024)
	for {
		n, _ := fl.Read(buf)
		if 0 == n {
			break
		}
		os.Stdout.Write(buf[:n])
	}
}


func main() {
	dirOp()

	fileOp()
}