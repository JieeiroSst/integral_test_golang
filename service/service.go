package service

import (
	"fmt"
	"integral_test/database"
)

type Service struct {
	Store database.Store
}

func (s *Service) GetNumber(ID int) error {
	result, err := s.Store.Get(ID)
	if err != nil {
		return err
	}
	if result > 10 {
		return fmt.Errorf("result too high: %d", result)
	}
	return nil
}

func NewGetNumber(store database.Store) func(int) error {
	return func(ID int) error {
		result, err := store.Get(ID)
		if err != nil {
			return err
		}
		if result > 10 {
			return fmt.Errorf("result too high: %d", result)
		}
		return nil
	}
}