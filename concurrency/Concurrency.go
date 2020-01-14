package main

import (
	"fmt"
	"time"
)

func main() {
	test01Timer()
	test02PingPong()
}

func test01Timer() {
	for i := 0; i < 24; i++ {
		c := timer(1 * time.Microsecond)
		fmt.Println(<-c)
	}
}

func timer(d time.Duration) <-chan int {

	c := make(chan int)
	go func() {
		time.Sleep(d)
		c <- 1
	}()
	return c

}

func test02PingPong() {

	var Ball int
	table := make(chan int)
	go player(table)
	go player(table)

	table <- Ball
	time.Sleep(1 * time.Second)
	<-table
}

func player(table chan int) {

	for {
		ball := <-table
		ball++
		fmt.Println(ball)
		time.Sleep(100 * time.Millisecond)
		table <- ball
	}

}
