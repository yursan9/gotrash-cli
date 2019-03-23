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
