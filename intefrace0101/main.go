package main

import "fmt"

type f func(c chan interface{})

func main() {

	fs := make([]f, 0)
	fs = append(fs, func(c chan interface{}) {
		s := <-c
		s2 := s.(string)
		c <- s2
	})

	out := make(chan interface{})
	for _, f := range fs {
		go f(out)
	}
	out <- "1"
	fmt.Println(<-out)

}
