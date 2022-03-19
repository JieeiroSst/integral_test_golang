package main

import (
	"github.com/stretchr/testify/mock"
	"testing"
)

type MockFilter struct {
	mock.Mock
}

func (m *MockFilter) Check(a int) bool {
	args := m.Called(a)
	return args.Bool(0)
}

func TestCalcMethod_Add(t *testing.T) {
	mf := new(MockFilter)
	mf.On("Check", 5).Return(false)

	clc := NewCalculatorWithFilter(mf)
	result := clc.Add(5, 9)

	t.Log(result)
	mf.AssertExpectations(t)
}