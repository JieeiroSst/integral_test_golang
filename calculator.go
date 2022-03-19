package main

type CalculatorInterface interface {
	Add(a int, b int) int
}

type Calculator struct {
	filter FilterInterface
}

func NewCalculatorWithFilter(filter FilterInterface) CalculatorInterface {
	return &Calculator{
		filter: filter,
	}
}

func (c *Calculator) Add(a int, b int) int {
	checked := c.filter.Check(a)
	if checked {
		return a + b
	} else {
		return 0
	}
}
