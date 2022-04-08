package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSum(t *testing.T) {
	assert.Equal(t, Sum(4, 6), 10)
}

func TestHello(t *testing.T) {
	res := Hello("Reza")
	assert.Equal(t, "Hello Reza", res)
}
