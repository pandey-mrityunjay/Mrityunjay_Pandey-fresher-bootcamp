package Testing2

import (
	"testing"
)

// Run test and debugs
func TestHelloWorld(t *testing.T) {
	result := HelloWorld("ADAM1")
	if result != "Hello ADAM" {
		t.Fatal("Result Wrong!")
	}
}

func TestHelloWorldJames(t *testing.T) {
	result := HelloWorld("ADAM")
	if result != "Hello ADAM" {
		t.Error("Result Wrong!")
	}
}
