package bob // package name must match the package name in bob_test.go

const testVersion = 2 // same as targetTestVersion

var bobResponses = map[string]string{
	"question":          "Sure",
	"yell":              "Whoa, chill out!",
	"nothing-to-saying": "Fine. Be that way!",
}
