package main

import "fmt"

// you can also use imports, for example:
// import "fmt"
// import "os"

// you can write to stdout for debugging purposes, e.g.
// fmt.Println("this is a debug message")

func Solution(A []int) int {
	// write your code in Go 1.4
	quickSort1(A, 0, len(A)-1)
	println("A：", A)
	key := 1
	for _, v := range A {
		if v <= 0 {
			continue
		}
		if v > key {
			return key
		}
		if v == key {
			key++
		}
	}
	return key
}

func quickSort1(arr []int, low, high int) {
	if low < high {
		i, j := low, high
		key := arr[(low+high)/2]
		for i <= j {
			for arr[i] < key {
				i++
			}
			for arr[j] > key {
				j--
			}
			if i <= j {
				arr[i], arr[j] = arr[j], arr[i]
				i++
				j--
			}
		}

		if low < j {
			quickSort1(arr, low, j)
		}
		if high > i {
			quickSort1(arr, i, high)
		}
	}

}

func main() {
	a := []int{1, 2, 32, 4, 5, 1, 2}
	quickSort1(a, 0, len(a)-1)
	fmt.Println("A:", a)
	b := Solution(a)
	println("b:", b)
}
