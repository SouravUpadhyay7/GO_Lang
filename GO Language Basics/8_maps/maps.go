package main

import "fmt"

func main() {
	var scores map[string]int = make(map[string]int)
	scores["Alice"] = 30
	scores["Bob"] = 25

	ages := map[string]int{"Charlie": 40, "Diana": 35}

	fmt.Println("Scores:", scores)
	fmt.Println("Ages:", ages)
}



// This program demonstrates the use of maps in Go.
// It declares a map to store scores of players and initializes it,
// then declares another map to store ages of people.