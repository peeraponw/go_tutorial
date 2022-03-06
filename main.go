package main

import "fmt"

func main() {
	// fmt.Println("Hello, World!")
	// fmt.Println(greetings("Putin"))
	// fmt.Println(greetingsWithAge("Putin", 45))
	// fmt.Println(greetingsWithAgeAndDrink("Putin", 45, "Vodka"))
	fmt.Println(sumOfFirst(3))
	fmt.Println(isPrime(1))
	fmt.Println(isPrime(2))
	fmt.Println(isPrime(3))
	fmt.Println(isPrime(4))
}

func greetings(name string) string {
	return "Hello " + name
}

func greetingsWithAge(name string, age uint) string {
	return "Hello " + name + ". You are " + fmt.Sprint(age) + " years old."
}

func greetingsWithAgeAndDrink(name string, age uint, drink string) string {
	return "Hello " + name + ". You are " + fmt.Sprint(age) + " years old. Your favorite drink is " + drink + "."
}

func sumOfFirst(n int) int {
	var s int = 0
	for i := 0; i <= n; i++ {
		s += i
	}
	return s
}

func isPrime(n int) bool {
	if n == 1 {
		return false
	}
	for i := 2; i < n; i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}
