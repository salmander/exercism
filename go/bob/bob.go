package bob // package name must match the package name in bob_test.go
import "strings"

const testVersion = 2 // same as targetTestVersion

var bobResponses = map[string]string{
	"question":       "Sure.",
	"yell":           "Whoa, chill out!",
	"nothing-to-say": "Fine. Be that way!",
	"anything-else":  "Whatever.",
}

func Hey(say string) string {
	var reply string

	// Remove any whitespace
	say = strings.TrimSpace(say)

	switch {
	// Check if someone has addressed him without actually saying anything
	// Fine. Be that way!
	case say == "":
		return bobResponses["nothing-to-say"]

	// Check if someone is yelling at Bob
	// Whoa, chill out!
	case say == strings.ToUpper(say) && say != strings.ToLower(say):
		return bobResponses["yell"]

	// Check if someone has asked a question to Bob
	// Sure.
	case strings.HasSuffix(say, "?") == true:
		return bobResponses["question"]

	// Anything else
	// Whatever.
	default:
		return bobResponses["anything-else"]
	}
	return reply
}
