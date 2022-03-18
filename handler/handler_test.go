package handler_test

import (
	"testing"

	"integral_test/handler"

	"github.com/stretchr/testify/assert"
)

func TestAdd(t *testing.T) {
	assert := assert.New(t)

	var tests = []struct {
		firstNumber  int
		secondNumber int
		expected     int
	}{
		{1, 2, 3},
		{-1, 1, 0},
		{1, -1, 2},
		{-1, -1, -2},
		{11111, 11111, 22222},
	}

	for _, test := range tests {
		assert.Equal(handler.Add(test.firstNumber, test.secondNumber), test.expected)
	}
}
