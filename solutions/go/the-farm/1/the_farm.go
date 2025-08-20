package thefarm

import (
	"errors"
	"fmt"
)

// TODO: define the 'DivideFood' function
func DivideFood(fc FodderCalculator, cowsNumber int) (float64, error) {
	totalFood, err := fc.FodderAmount(cowsNumber)
	if err != nil {
		return 0, err
	}
	fatteningFactor, err := fc.FatteningFactor()
	if err != nil {
		return 0, err
	}

	return totalFood * fatteningFactor / float64(cowsNumber), nil
}

// TODO: define the 'ValidateInputAndDivideFood' function
func ValidateInputAndDivideFood(fc FodderCalculator, cowsNumber int) (float64, error) {
	if cowsNumber <= 0 {
		return 0, errors.New("invalid number of cows")
	}
	return DivideFood(fc, cowsNumber)
}

// TODO: define the 'ValidateNumberOfCows' function

type InvalidCowsError struct {
	numberOfCows int
	message      string
}

func (e *InvalidCowsError) Error() string {
	return fmt.Sprintf("%d cows are invalid: %s", e.numberOfCows, e.message)
}

func ValidateNumberOfCows(cowsNumber int) error {
	if cowsNumber < 0 {
		return &InvalidCowsError{
			numberOfCows: cowsNumber,
			message:      "there are no negative cows",
		}
	}
	if cowsNumber == 0 {
		return &InvalidCowsError{
			numberOfCows: cowsNumber,
			message:      "no cows don't need food",
		}
	}
	return nil
}

// Your first steps could be to read through the tasks, and create
// these functions with their correct parameter lists and return types.
// The function body only needs to contain `panic("")`.
//
// This will make the tests compile, but they will fail.
// You can then implement the function logic one by one and see
// an increasing number of tests passing as you implement more
// functionality.
