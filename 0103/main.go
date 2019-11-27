package main

import "fmt"

func main() {
	s, i := some("s", 1)
	fmt.Println(s, i)
}

func some(s string, i int) (s1 string, i1 int)  {

	//var s1 string = "s1"

	s1 = "s1"
	i1 = 1

	s1 += s
	i1 += i

	return
}
