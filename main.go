package main

import "fmt"

func main() {
	var name string
	// var time uint
	// var break bool

	fmt.Println("What is your name ?")
	fmt.Scan(&name)
	fmt.Printf("Welcome %v \n", name)
}
