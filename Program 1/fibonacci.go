package main

/**
  Recursive Fibonacci program to calculate the nth Fibonacci number. f[n-1] + f[n-2]
**/
import (
	"fmt"
	"time"
)

/* Recursive Fibonacci function */
func fib(n int) int {
	if n == 0 {
		return 0
	}
	if n == 1 {
		return 1
	}
	return fib(n-1) + fib(n-2)
}

/* Main routine to run Fibonacci program */
func main() {
	fmt.Print("* * * Fibonacci Printer * * *\n\n")
	fmt.Print("Which Fibonacci number would you like to see?: ")
	var fibNum int
	fmt.Scanln(&fibNum)

	if fibNum > 0 && fibNum < 46 {
		timeStart := time.Now()
		result := fib(fibNum)
		timeElapsed := time.Since(timeStart)

		fmt.Printf("\nFibonacci number %d is: %d\n\n", fibNum, result)
		fmt.Printf("This calculation required %.4v seconds\n\n", timeElapsed.Seconds())
	} else {
		fmt.Println("\nError: entry must be from 1 to 45 inclusive.")
	}
}
