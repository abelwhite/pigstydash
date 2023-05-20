package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// Use case: 1 Register
func TestIsEmailValid(t *testing.T) {
	assert.Equal(t, true, isEmailValid("test44@gmail.com"))
	assert.Equal(t, true, isEmailValid("shouldwork@gmail.com"))
	assert.Equal(t, false, isEmailValid("test44$@gmail.com"))
}

func TestIsPasswordValid(t *testing.T) {
	assert.Equal(t, true, isPassValid("P@ssw0rd"))
	assert.Equal(t, false, isPassValid("pass"))
}


