package helper

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

// //* Fail()
// func TestHelloWorldEko(t *testing.T) {
// 	result := HelloWorld("jery")
// 	if result != "Hello, jery" {
// 		// error
// 		t.Fail()
// 	}
// 	fmt.Println("TestHelloWorldEko Done")
// }

// //*  FailNow()
// func TestHelloWorldJery1(t *testing.T) {
// 	result := HelloWorld("jery")
// 	if result != "Hello, jery" {
// 		// error
// 		t.FailNow()
// 	}
// 	fmt.Println("TestHelloWorldJery Done")
// }

// //* t.Error()
// func TestHelloWorldEko1(t *testing.T) {
// 	result := HelloWorld("jery")
// 	if result != "Hello, jery" {
// 		// error
// 		t.Error("result is not Hello, jery")
// 	}
// 	fmt.Println("TestHelloWorldEko Done")
// }

// //* t.Fatal
// func TestHelloWorldJery(t *testing.T) {
// 	result := HelloWorld("jery")
// 	if result != "Hello, jery" {
// 		// error
// 		t.Fatal("result is not Hello, jery")
// 	}
// 	fmt.Println("TestHelloWorldJery Done")
// }

// * Test Assertion with Testify
func TestHelloWorldAssert(t *testing.T) {
	result := HelloWorld("jery")
	assert.Equal(t, "Hello, jery", result, "Result must be 'Hello, jery'")
	fmt.Println("TestHelloWorldAssert Done")
}
