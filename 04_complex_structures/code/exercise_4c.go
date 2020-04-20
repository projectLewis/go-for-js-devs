package main

import (
	"fmt"
	"strings"
)

func averageScore(scores []float64) (average float64) {
	var total float64
	for _, score := range scores {
		total += score
	}
	average = total / float64(len(scores))
	return
}

func nameExists(name string, list map[string]string) bool {
	nameTitle := strings.ToLower(name)
	_, ok := list[nameTitle]
	return ok
}

func addGroceries(groceryList []string, groceryItem ...string) []string {
	updatedList := append(groceryList, groceryItem...)
	return updatedList
}

var petNames = map[string]string{
	"kima":   "dog",
	"stella": "dog",
	"batman": "cat",
}

func main() {
	scores := []float64{95.9, 84.2, 99.0, 100.0, 22.3}
	fmt.Println(averageScore(scores))

	fmt.Println(nameExists("kiMa", petNames))
	fmt.Println(nameExists("bond", petNames))

	groceries := []string{"yams", "cakes", "potatoes", "napkins", "spoons"}
	groceryList := groceries[2:5]
	fmt.Println(groceryList)
	updatedList := addGroceries(groceryList, "chocolate", "forks")
	fmt.Println(updatedList)
}
