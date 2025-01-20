package main

import (
	"fmt"
	"time"
)

func main() {
	var birthYear int

	fmt.Println("\t\t\tAGE CALCULATOR\t\t\t")
	fmt.Println("\t\t\t===============\t\t\t\n")
	fmt.Print("Enter your year of birth: ")
	fmt.Scanln(&birthYear)

	currentYear := time.Now().Year()
	age := currentYear - birthYear
	fmt.Printf("You are %d years old", age)
}
