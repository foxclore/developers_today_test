package models

import (
	"developers_today_test/breeds"
	"errors"
)

type Cat struct {
	Name       string  `json:"name"`
	Experience int     `json:"experience"`
	Breed      string  `json:"breed"`
	Salary     float64 `json:"salary"`
}

func (c *Cat) ValidateBreed() error {
	if !breeds.ValidateBreed(c.Breed) {
		return errors.New("breed doesn't exist")
	}
	return nil
}

func (c *Cat) Validate() error {
	if c.Name == "" {
		return errors.New("empty name")
	}
	return c.ValidateBreed()
}
