package main

import (
	"fmt"
	"strconv"
	"sync"
	"time"
)

type job func(in, out chan interface{})

func single(in, out chan interface{}) {

	var s string

	for i := 0; i <= len(in); i++ {
		i := <-in
		s = i.(string)
		s += "~" + s
		out <- s
	}

}

func multi(in, out chan interface{}) {

	var s string

	for i := 0; i <= len(in); i++ {
		var result string
		i := <-in
		s = i.(string)

		for j := 0; j <= 5; j++ {
			result += strconv.Itoa(j) + "+" + s
		}

		out <- result
	}

}

func combine(in, out chan interface{}) {

	var s string
	for i := 0; i <= len(in); i++ {
		i := <-in
		s = i.(string)
		s += "_" + s
	}
	out <- s
}

func execute(jobs ...job) {

	in := make(chan interface{}, 7)

	wg := &sync.WaitGroup{}
	wg.Add(len(jobs))
	for _, j := range jobs {
		out := make(chan interface{}, 7)
		go func(out chan interface{}) {
			j(in, out)
			close(out)
			wg.Done()
		}(out)
		for i := range out {
			fmt.Println(i)
			in <- i
		}
	}
	wg.Wait()
	close(in)

}

func main() {

	inputData := []int{0, 1, 1, 2, 3, 5, 8}

	jobs := []job{
		job(func(in, out chan interface{}) {
			for _, i := range inputData {
				out <- strconv.Itoa(i)
			}
		}),
		job(single),
		job(multi),
		job(combine),
		//job(p),
	}

	execute(jobs...)

	time.Sleep(1000 * time.Millisecond)

}
