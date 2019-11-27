package main

import "fmt"

type Vehicle interface {
	save()
}

type Bmw struct {
	name string
}

type Mercedes struct {
	name string
}

func (v Bmw) save() {
	fmt.Println("this is: " + v.name)
}

func (v Mercedes) save() {
	fmt.Println("this is: " + v.name)
}

func store (v Vehicle) {
	fmt.Println(v.save)
}

func main() {

	//b := Bmw{name: "bmw"}
	//m := Mercedes{name: "mercedes"}
	//
	//vehicles := []Vehicle{b, m}
	//for _, v := range vehicles {
	//	store(v)
	//}
	i := make([]int, 0) // var i []int
	fmt.Println(len(i))
	i = append(i, 0)
	fmt.Println(len(i))
}
