package greet

import "testing"

func TestHello(t *testing.T) {
	got := hello()
	want := "Hello, World"

	if got != want {
		t.Errorf("Got %q, want %q\n", got, want)
	}
}

func TestHelloName(t *testing.T) {
	tests := []struct {
		got  string
		want string
		name string
	}{
		{got: helloName("pt", "Jhon"), want: "Oi, Jhon", name: "Portuguese"},
		{got: helloName("en", "Alice"), want: "Hello, Alice", name: "English"},
		{got: helloName("es", "Marcos"), want: "Hola, Marcos", name: "Spanish"},
		{got: helloName("fr", "Marcos"), want: "Bonjour, Marcos", name: "French"},
		{got: helloName("", "Hans"), want: "?, Hans", name: "Unknown Language"},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			if test.got != test.want {
				t.Errorf("Got %q, want %q\n", test.got, test.want)
			}
		})
	}
}
