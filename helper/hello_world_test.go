package helper

import (
	"fmt"
	"runtime"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
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

// * With Assert
func TestHelloWorldAssert(t *testing.T) {
	result := HelloWorld("jery")
	assert.Equal(t, "Hello, jery", result, "Result must be 'Hello, jery'")
	fmt.Println("TestHelloWorldAssert Done")
}

//* With Require
func TestHelloWorldRequire(t *testing.T) {
	result := HelloWorld("jery")
	require.Equal(t, "Hello, jery", result, "Result must be 'Hello, jery'")
	fmt.Println("TestHelloWorldAssert Done")
}

//  * Skip Test
func TestSkip(t *testing.T) {
	if runtime.GOOS == "windows" {
		t.Skip("Can not run on windows")
	}

	result := HelloWorld("jery")
	require.Equal(t, "Hello, jery", result, "Result must be 'Hello, jery'")
}

//* Before dan After Test
func TestMain(m *testing.M) {
	fmt.Println("TestBefore Done")
	m.Run()
	fmt.Println("TestAfter Done")
}

//* Sub Test
func TestSubTest(t *testing.T) {
	t.Run("TestSubTest1", func(t *testing.T) {
		result := HelloWorld("jery")
		require.Equal(t, "Hello, jery", result, "Result must be 'Hello, jery'")
	
	})

	t.Run("TestSubTest2", func(t *testing.T) {
		result := HelloWorld("eko")
		require.Equal(t, "Hello, eko", result, "Result must be 'Hello, eko'")
	
	})
}

//* Table Test
func TestHelloWorldTable(t *testing.T) {
	var tests = []struct {
		name string
		request   string
		expected string
	}{
		{name: "jery", request: "jery", expected: "Hello, jery"},
		{name: "eko", request: "eko", expected: "Hello, eko"},
		{name: "budi", request: "budi", expected: "Hello, budi"},
		{name: "tomo", request: "tomo", expected: "Hello, tomo"},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := HelloWorld(test.request)
			require.Equal(t, test.expected, result)
		})
	}
}














