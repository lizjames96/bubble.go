package main

import "fmt"

func BubbleSort(a []int) {
	if a == nil {
		return
	}
	n := len(a)
	for i := n; i > 0; i-- {
		for j := 0; j+1 < i; j++ {
			if a[j] > a[j+1] {
				Swap(a, j)
			}
		}
	}
}

func Swap(a []int, i int) {
	if a == nil || i < 0 || i+1 > len(a)-1 {
		// null check and index check
		return
	}
	tmp := a[i]
	a[i] = a[i+1]
	a[i+1] = tmp
}

func InputSlice() []int {
	var a = make([]int, 0, 3)
	fmt.Printf("Input number of integers: ")

	var n int
	var val int
	fmt.Scan(&n)
	fmt.Printf("Please input %d integers.\n", n)
	for i := 0; i < n; i++ {
		fmt.Scan(&val)
		a = append(a, val)
	}
	return a
}

func main() {
	a := InputSlice()
	BubbleSort(a)
	fmt.Printf("%v\n", a)
}