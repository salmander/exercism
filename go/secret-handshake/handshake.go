package secret

import "fmt"

var p = fmt.Println

func Handshake(number int) []string {
	var returnString []string

	if number <= 0 {
		return returnString
	}

	if number&1 != 0 {
		returnString = append(returnString, "wink")
	}

	if number&2 != 0 {
		returnString = append(returnString, "double blink")
	}

	if number&4 != 0 {
		returnString = append(returnString, "close your eyes")
	}

	if number&8 != 0 {
		returnString = append(returnString, "jump")
	}

	if number&16 != 0 {
		for i, j := 0, len(returnString)-1; i < j; i, j = i+1, j-1 {
			returnString[i], returnString[j] = returnString[j], returnString[i]
		}
	}

	return returnString
}
