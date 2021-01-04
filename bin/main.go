package main

import "fmt"

func maxDivide(a int, b int) int {
	for a%b == 0 {
		a = a / b
	}
	return a
}

func isUgly(no int) int {
	no = maxDivide(no, 2)
	no = maxDivide(no, 3)
	no = maxDivide(no, 5)

	if no == 1 {
		return 1
	}
	return 0
}

func getNthUglyNo(n int) int {
	i := 1

	// ugly number count
	count := 1

	// check for all integers
	// until count becomes n
	for n > count {
		i++

		if isUgly(i) == 1 {
			count++
		}

	}
	return i
}

func main() {
	i := 10
	fmt.Printf("[INPUT] %d\n", i)
	o := getNthUglyNo(i)
	fmt.Printf("[OUTPUT] %d\n", o)
}
