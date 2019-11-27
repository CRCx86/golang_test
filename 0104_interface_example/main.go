package main

import "fmt"

type A interface {
	write()
	read()
}

type B struct {
	b string
}

func readWrite(ab A) {
	ab.read()
	ab.write()
}

func (b *B) write() {
	fmt.Println("writing")
}

func (b *B) read() {
	fmt.Println("reading")
}

func main() {

	c := new(B)
	readWrite(c)

	d := B{}
	readWrite(&d)

	e := &B{}
	readWrite(e)

	letters := [...]A{c, &d, e}
	for index, letter := range letters{
		fmt.Println(index)
		letter.read()
		letter.write()
	}

}
