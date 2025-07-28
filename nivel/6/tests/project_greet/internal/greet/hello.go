package greet

import "fmt"

func hello() string {
	return "Hello, World"
}

func helloName(lang, name string) string {
	if lang == "pt" {
		return fmt.Sprintf("Oi, %s", name)
	}
	if lang == "es" {
		return fmt.Sprintf("Hola, %s", name)
	}
	if lang == "fr" {
		return fmt.Sprintf("Bonjour, %s", name)
	}
	if lang == "" {
		return fmt.Sprintf("?, %s", name)
	}

	return fmt.Sprintf("Hello, %s", name)
}
