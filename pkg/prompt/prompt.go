package prompt

import (
	"fmt"
	"strings"
)

func Number(text string) int {
	var answer int

	fmt.Printf("%s? ", text)
	fmt.Scanln(&answer)

	return answer
}

func YesNo(text string) bool {
	var answer string

	fmt.Printf("%s [Yes/No]? ", text)
	fmt.Scanln(&answer)

	if strings.HasPrefix(answer, "y") || strings.HasPrefix(answer, "Y") {
		return true
	}

	return false
}
