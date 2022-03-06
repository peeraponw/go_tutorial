package main

import "fmt"

func main() {
	fmt.Println("Hello, World!")
	fmt.Println(greetings("Putin"))
	fmt.Println(greetingsWithAge("Putin", 45))
	fmt.Println(greetingsWithAgeAndDrink("Putin", 45, "Vodka"))
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
