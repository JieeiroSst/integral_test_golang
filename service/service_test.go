package service_test

import (
	"errors"
	"integral_test/database"
	"integral_test/service"
	"testing"
)

func TestServiceSuccess(t *testing.T) {
	m := new(database.MockStore)
	m.On("Get", 2).Return(7, nil)
	s := service.Service{m}
	err := s.GetNumber(2)
	m.AssertExpectations(t)
	if err != nil {
		t.Errorf("error should be nil, got: %v", err)
	}
}

func TestServiceResultTooHigh(t *testing.T) {
	m := new(database.MockStore)
	m.On("Get", 2).Return(24, nil)
	s := service.Service{m}
	err := s.GetNumber(2)
	m.AssertExpectations(t)
	if err.Error() != "result too high: 24" {
		t.Errorf("error should be 'result too high: 24', got: %v", err)
	}
}

func TestServiceStoreError(t *testing.T) {
	m := new(database.MockStore)
	m.On("Get", 2).Return(0, errors.New("failed"))
	s := service.Service{m}
	err := s.GetNumber(2)
	m.AssertExpectations(t)
	if err.Error() != "failed" {
		t.Errorf("error should be 'failed', got: %v", err)
	}
}
