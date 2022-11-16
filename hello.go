// i followed this tutorial on github - https://github.com/quii/learn-go-with-tests/blob/main/hello-world.md
package main

import "fmt"

const spanish = "Spanish"
const french = "French"
const englishHelloPrefix = "Hello "
const spanishHelloPrefix = "Hola "
const frenchHelloPrefix = "Bonjour "

// functions that starts lowercase are private
// functions that starts uppercase are public

func Hello(name string, language string) string {

	if name == "" {
		name = "World!"
	}

	return greetingPrefix(language) + name
}

func greetingPrefix(language string) (prefix string) { // with (prefix string) we say that return should return prefix. So we don't need return prefix
	switch language {
	case french:
		prefix = frenchHelloPrefix
	case spanish:
		prefix = spanishHelloPrefix
	default:
		prefix = englishHelloPrefix
	}
	return
}

func main() {
	fmt.Println(Hello("world", ""))
}
