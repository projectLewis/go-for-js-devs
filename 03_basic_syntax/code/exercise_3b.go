package main

import "fmt"

// func scanString(statement string) string {
// 	var reference string
// 	fmt.Println(statement)
// 	if _, err := fmt.Scanln(&reference); err != nil {
// 		fmt.Println(err)
// 	}
// 	return reference
// }
// func scanInt(statement string) int {
// 	var reference int
// 	fmt.Println(statement)
// 	if _, err := fmt.Scanln(&reference); err != nil {
// 		fmt.Println(err)
// 	}
// 	return reference
// }
// func scanBool(statement string) bool {
// 	var reference bool
// 	fmt.Println(statement)
// 	if _, err := fmt.Scanln(&reference); err != nil {
// 		fmt.Println(err)
// 	}
// 	return reference
// }

func getInput(prompt string, nval *string) {
	fmt.Println(prompt)
	if _, err := fmt.Scanln(nval); err != nil {
		fmt.Println(err)
	}
}

func main() {
	var name, region string
	getInput("Enter a name", &name)
	getInput("Enter a region", &region)
	fmt.Printf("Hi! My name is %s, and I'm from %s", name, region)
}
