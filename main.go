package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	var name string
	var quick_break string
	rand.Seed(time.Now().UnixNano())
	options := map[string][]string{
		"yes": {"Go watch some TikToks",
			"Watch a Youtube video",
			"Refill your water and take a walk",
			"Find someone to bother for a few minutes",
			"Grab a quick snack",
			"Go juggle a soccer ball for a bit",
			"Go back to work, sorry no break for you",
			"Take some deep breathes and recenter yourself"},
		"no": {"Catch up on soccer highlights",
			"Watch some anime",
			"Put on a docutmentary",
			"Play some video games",
			"Learn about a new coding language/framework",
			"Refactor code from previous projects",
			"Start a new project",
			"Reach out to someone you miss",
			"Take a nap"},
	}
	fmt.Println("What is your name ?")
	fmt.Scan(&name)
	fmt.Printf("Welcome %v, let's help you choose an activity \n", name)
	fmt.Println("Is this a quick break ? (yes or no) ")
	fmt.Scan(&quick_break)
	activity := options[quick_break][rand.Intn(7)]
	fmt.Printf("To defeat your boredom you should ... \n %v\n", activity)
}
