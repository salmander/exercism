package hello

const testVersion = 2

func HelloWorld(word string) string {
	if len(word) < 1 {
		word = "World"
	}

	return "Hello, " + word + "!"

}
