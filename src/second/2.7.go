package main

import (
	"fmt"
	"runtime"
	"time"
)

func say(s string) {
	for i := 0; i<5; i++ {
		runtime.Gosched()
		fmt.Println(s)
	}
}

func sum(a []int, c chan int) {
	total := 0
	for _, v := range a{
		total += v
	}
	fmt.Println(total)
	c <- total
}

func write2Chan() {
	c := make(chan int, 2)
	c <- 1
	c <- 2
	fmt.Println(<-c)
	fmt.Println(<-c)
	c <- 3
	fmt.Println(<-c)
}

func fibonacci(n int, c chan int) {
	x, y := 1, 1
	for i := 0; i<n; i ++ {
		c <- x
		x, y = y, x+y
	}
	close(c)
}

func fmtFiBo() {
	c := make(chan int, 10)
	go fibonacci(cap(c), c)
	for i:=range c{
		fmt.Println(i)
	}
}

func selectFb(c, q chan int) {
	x, y := 1, 1
	for {
		select {
		case c <- x:
			x, y = y, x+y
		case <- q:
			fmt.Println("Quit")
			return
		default:
			time.Sleep(1*time.Second)
		}
	}
}

func selectFbMain() {
	c := make(chan int)
	quit := make(chan int)
	go func() {
		for i:=0; i<10; i++{
			fmt.Println(<-c)
		}
		quit <- 0
	}()
	selectFb(c, quit)
}


func main() {
	go say("world")
	say("Hello")

	a := []int{1, 3, 6, 9, 11, 20}
	c := make(chan int, 5)
	go sum(a[:len(a)/2], c)
	go sum(a[len(a)/2:], c)
	x, y := <-c, <-c
	fmt.Println(x, y, x+y)
	write2Chan()

	fmtFiBo()

	selectFbMain()

}
