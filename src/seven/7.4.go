package main

import (
	"fmt"
	"time"
)

func main() {

	t := time.Now()
	for i :=0 ; i < 100000 ; i++ {
		for j := 0; j < 100; j++{

		}
	}
	t1 := time.Now()

	dt := t1.Nanosecond() - t.Nanosecond()
	fmt.Println("外大内小", dt)


	t2 := time.Now()
	for i :=0 ; i < 100; i++ {
		for j := 0; j < 100000; j++{

		}
	}
	t3 := time.Now()

	dt1 := t3.Nanosecond() - t2.Nanosecond()
	fmt.Println("外小内大", dt1)
}
