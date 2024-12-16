package main

import "fmt"

func factorialLoop(value int) int {
	result := 1
	for i := value; i > 0; i-- {
		result *= i
		fmt.Println(result)
	}
	return result
}

func recursiveFactorial(value int) int {
	if value == 1 {
		return 1
	} else {
		return value * recursiveFactorial(value-1)
	}
}

func main() {
	fmt.Println(factorialLoop(10))
	fmt.Println("++++++ ",recursiveFactorial(10))
}