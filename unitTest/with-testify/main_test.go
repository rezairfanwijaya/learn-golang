package main

import (
	"fmt"
	"runtime"
	"testing"

	"github.com/stretchr/testify/assert"
)

// test main untuk implementasi before dan after testing
func TestMain(m *testing.M) {
	// before
	fmt.Println("Before testing")

	// run all test
	m.Run()

	// after testing
	fmt.Println("After testing")
}

func TestSum(t *testing.T) {
	assert.Equal(t, Sum(4, 6), 10)
}

func TestHello(t *testing.T) {
	res := Hello("Reza")
	assert.Equal(t, "Hello Reza", res)
}

func TestSkip(t *testing.T) {
	if runtime.GOOS == "windows" {
		t.Skip("This test does not run on windows")
	}

	fmt.Println("Ini testing ")
}

// implementasi subtest
func TestBagi(t *testing.T) {
	t.Run("Not_nol", func(t *testing.T) {
		assert.Equal(t, bagi(10, 2), 5)
	})

	t.Run("With_nol", func(t *testing.T) {
		assert.Equal(t, bagi(10, 0), "Nol tidak boleh")
	})
}
