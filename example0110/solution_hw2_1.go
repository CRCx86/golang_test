package main

import (
	"fmt"
	"strconv"
)

type jobb func(out chan interface{})

func main() {

	inputData := []int{0, 1}

	jobs := []jobb{
		jobb(func(out chan interface{}) {
			for _, i := range inputData {
				out <- strconv.Itoa(i)
			}
		}),
		jobb(func(out chan interface{}) {
			for o := range out {
				fmt.Println(o)
			}
			close(out)
		}),
	}

	for _, j := range jobs {
		out := make(chan interface{})
		go j(out)
	}

}
