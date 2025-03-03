package Testing3

import (
	"fmt"
	Testing2 "golang-testing/Golang-Day-3/Testing/Unit-Testing-2"
	"testing"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestHelloWorld(t *testing.T) {
	result := Testing2.HelloWorld("Adamm")
	assert.Equal(t, "Hello Adam", result, "Wrong output:")
	fmt.Println("Hello World Test 3 Done:")

}

func TestHelloWorld2(t *testing.T) {
	fmt.Println("Hello World Test Using Require")
	result := Testing2.HelloWorld("Adam")
	require.Equal(t, "Hello Adam", result, "Wrong Output")

}
