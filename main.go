package main

import "fmt"

func main() {
	var name string
	var time uint
	var quick_break string

	fmt.Println("What is your name ?")
	fmt.Scan(&name)
	fmt.Printf("Welcome %v \n", name)
	fmt.Println("How much time do you have (in min)?")
	fmt.Scan(&time)
	fmt.Println("Is this a break yes or no ?")
	fmt.Scan(&quick_break)

}
