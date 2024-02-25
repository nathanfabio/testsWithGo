package main

import "fmt"

const prefixPt = "Olá, "
const prefixEn = "Hello, "
const prefixFr = "Bonjour, "
// const portuguese = "Portuguese"
// const french = "French"

func Hello(name string, language string) string {
	if name == "" {
		name = "World"
	}

	return greentingPrefix(language) + name
}

func greentingPrefix(language string) (prefix string) {
	switch language {
	case "French":
		prefix = prefixFr
	case "Portuguese":
		prefix = prefixPt
	default:
		prefix = prefixEn
	}
	return
}

func main() {
	fmt.Println(Hello("World" , ""))
}