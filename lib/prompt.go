package lib

import "fmt"

func prompt(text string) bool {
	var answer string

	fmt.Print(text, " [Yes/No]? ")
	fmt.Scanln(&answer)

	if answer == "y" || answer == "Y" || answer == "yes" || answer == "Yes" {
		return true
	}

	return false
}

func promptInt(text string) int {
	var answer int
	
	fmt.Print(text, "? ")
	fmt.Scanln(&answer)
	
	return answer
}
