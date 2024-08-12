package main

import "fmt"

func fib2(n int) int {
	if n <= 1 {
		return n
	}
	f := make([]int, n+2)
	f[0], f[1] = 0, 1
	for i := 2; i <= n; i++ {
		f[i] = f[i-1] + f[i-2]
	}
	return f[n]
}

func main() {
	var calcFib = fib2(100)
	fmt.Printf("Test %s is %d \n", "fibonachi", calcFib)

}
