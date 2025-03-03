package Testing1

import (
	"testing"
)

// Run test and debugs
func TestHelloWorld(t *testing.T) {
	result := HelloWorld("ADAM")
	if result != "Hello ADAM" {
		panic("Result Wrong!")
	}
}
