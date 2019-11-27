package main

import (
	"fmt"
	"math"
)

func main() {

	var a int
	var b int

	fmt.Scan(&a)
	fmt.Scan(&b)

	a1 := make([]int, int(math.Floor(math.Log10(float64(a)))) + 1)
	b1 := make([]int, int(math.Floor(math.Log10(float64(b)))) + 1)

	a1 = initArray(a, a1)
	b1 = initArray(b, b1)

	a1 = reverse(a1)
	b1 = reverse(b1)

	for _, ai := range a1{
		if ai == 0 {
			break
		}
		for _, bj := range b1 {
			if bj == 0 {
				break
			}
			if ai == bj {
				fmt.Println(ai)
			}
		}
	}

}

func initArray(n int, m []int) []int {

	for i := 0; n != 0; i++ {
		b := n % 10
		n = n / 10
		m[i] = b
	}

	return m
}

func reverse(m []int) []int {

	for left, right := 0, len(m) - 1; left < right; left, right = left + 1, right - 1 {
		m[left], m[right] = m[right], m[left]
	}
	
	return m
}
