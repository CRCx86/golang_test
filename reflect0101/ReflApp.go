package main

import (
	"fmt"
	"reflect"
)

type T struct {
	A int
	B string
}

func (t *T) Save() (r int){
	r = 123
	return
}

func main() {

	//Example 1
	var f float64
	fmt.Println("value f is", reflect.TypeOf(f))
	fmt.Println("value f is", reflect.ValueOf(f).String())
	w := reflect.ValueOf(f)
	y := w.Interface().(float64)
	fmt.Println("value of y", y)

	// Example 2
	var x float64 = 3.4
	p := reflect.ValueOf(&x)
	fmt.Println("type of p: ", p.Type())
	fmt.Println("settability p: ", p.CanSet())
	v := p.Elem()
	fmt.Println("settability v: ", v.CanSet())
	v.SetFloat(7.1)
	fmt.Println(v.Interface())
	fmt.Println(x)

	// Example 3
	fmt.Println("Fields")
	t := T{23, "skidoo"}
	s := reflect.ValueOf(&t).Elem()
	typeOfT := s.Type()
	for i := 0; i < s.NumField(); i++ {
		f := s.Field(i)
		fmt.Printf("%d: %s %s = %v\n", i, typeOfT.Field(i).Name, f.Type(), f.Interface())
	}

	s.Field(0).SetInt(77)
	s.Field(1).SetString("Sunset Strip")
	fmt.Println("t is now", t)

	fmt.Println("type of t is: ", reflect.TypeOf(t))


	m := reflect.TypeOf(&t)
	fmt.Println("Methods")
	for i := 0; i < m.NumMethod(); i++  {
		f := m.Method(i)
		fmt.Printf("%s\n", f.Name)
	}

	m = reflect.TypeOf(&t)
	fmt.Printf("Methods %v\n", m.NumMethod())
	for i := 0; i < m.NumMethod(); i++  {
		//f := m.Method(i)
		fmt.Printf("%s\n", m.Method(i).Name)
	}

	fmt.Printf("%s", reflect.ValueOf(&t).MethodByName("Save"))

}
