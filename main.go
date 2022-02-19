package main

import "fmt"

func main() {
	var name string
	var quick_break string
	options := map[string][]string{
		"yes": {},
		"no":  {},
	}
	fmt.Println("What is your name ?")
	fmt.Scan(&name)
	fmt.Printf("Welcome %v \n", name)
	fmt.Println("Is this a break yes or no ?")
	fmt.Scan(&quick_break)

}
