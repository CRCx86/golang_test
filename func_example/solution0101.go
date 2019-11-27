package main

import "fmt"

func main() {

	for i := 0; i < 100; i++ {
		f := fun1()
		go f(i, i + 1)
	}

}

func fun1() func(s, i int) {

	return func(s, i int) {
		fmt.Println(s + i)
	}

}