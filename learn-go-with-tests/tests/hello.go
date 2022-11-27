package tests

import "fmt"

const (
	english Lang = "English"
	spanish Lang = "Spanish"
	french  Lang = "French"
	german  Lang = "German"
)

var langPrefix = map[Lang]string{
	english: "Hello",
	spanish: "Hola",
	french:  "France",
	german:  "Bonjour",
}

// Lang type of language
type Lang string

// Hello say welcome to go
func Hello(name, lang string) string {
	if name == "" {
		name = "World"
	}

	prefix, ok := langPrefix[Lang(lang)]
	if !ok {
		prefix = langPrefix[english]
	}

	return fmt.Sprintf("%s, %s", prefix, name)
}

// Add return two int sum value
func Add(x, y int) int {
	return x + y
}

// Repeat generates a duplicated chars
func Repeat(char string, num int) string {
	result := ""
	for i := 0; i < num; i++ {
		result += char
	}
	return result
}
