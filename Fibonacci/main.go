package main

import "fmt"

func fibonacci(n int) int {
	if n <= 1 {
		return n
	}
	return fibonacci(n-1) + fibonacci(n-2)
}

func main() {

	var t int
	fmt.Printf("Sequência de Fibonacci.\nDigite quantos termos devem aparecer: ")
	fmt.Scanf("%d", &t)

	for i := 0; i < t; i++ {
		fmt.Printf("%v ", fibonacci(i))
	}

}
