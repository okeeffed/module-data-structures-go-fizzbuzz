package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFizz(t *testing.T) {
	res := fizzBuzz(3)
	if res != "fizz" {
		t.Errorf("Sum was incorrect, got: %s, want: %s.", res, "fizz")
	}
}

func TestBuzz(t *testing.T) {
	res := fizzBuzz(5)
	if res != "buzz" {
		t.Errorf("Sum was incorrect, got: %s, want: %s.", res, "buzz")
	}
}

func TestFizzBuzz(t *testing.T) {
	res := fizzBuzz(15)
	if res != "fizzbuzz" {
		t.Errorf("Sum was incorrect, got: %s, want: %s.", res, "fizzbuzz")
	}
}

func TestInt(t *testing.T) {
	res := fizzBuzz(7)
	if res != "7" {
		t.Errorf("Sum was incorrect, got: %s, want: %s.", res, "7")
	}
}

/**
 * Using the testify library
 */
func TestUsingTestify(t *testing.T) {
	assert := assert.New(t)
	res := fizzBuzz(7)
	assert.Equal(res, "7", "Non-fizzbuzz result works")
}
