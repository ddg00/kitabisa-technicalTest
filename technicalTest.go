package main

import (
	"fmt"
	"math"
	"strconv"
)
// Function Sum
func sum(x int, y int) int {
	return x + y
}
// Function Multiply
func time(x int, y int ) int {
	return x * y
}

func isPrime(value int) bool {
	for i := 2; i <= int(math.Floor(float64(value)/2)); i++ {
		if value%i == 0 {
			return false
		}
	}
	return value > 1
}

// Fucntion get prime number
func getPrime(n int) {
	var nPrime []int;
	i := 0;
	for len(nPrime) <= n {
		if isPrime(i) {
			nPrime = append(nPrime, i);
		}
		i++;
	}
	for n:= 0; n < len(nPrime); n++ {
		fmt.Print(strconv.Itoa(nPrime[n])+ " ");
	};
}

func FibonacciRecursion(n int) int {
	if n <= 1 {
		return n
	}
	return FibonacciRecursion(n-1) + FibonacciRecursion(n-2)
}

// Fucntion get fibonancy sequence
func getFibonancySequence(n int){
	if n > 0 {
		for i := 0; i < n; i++ {
			fmt.Print(strconv.Itoa(FibonacciRecursion(i)) + " ")
		}
	}
}

func main() {

	fmt.Println(sum(1,2));
	fmt.Println(time(1,2));
	getPrime(4);
	fmt.Println(" ");
	getFibonancySequence(9);
}




