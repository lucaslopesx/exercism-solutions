package thefarm

import (
	"errors"
	"fmt"
)

var ErrInvalidNumberOfCows = errors.New("invalid number of cows")

type InvalidCowsError struct {
	numberOfCows int
	message string
}

func (err *InvalidCowsError) Error() string {
	return fmt.Sprintf("%d cows are invalid: %s", err.numberOfCows, err.message)
}


// TODO: define the 'DivideFood' function
func DivideFood(fodderCalculator FodderCalculator, numberOfCows int) (float64, error) {
	if numberOfCows <= 0 {
		return 0, ErrInvalidNumberOfCows
	}

	fodderAmount, err := fodderCalculator.FodderAmount(int(numberOfCows))
	if err != nil {
		return 0, err
	}

	fatteningFactor, err := fodderCalculator.FatteningFactor()
	if err != nil {
		return 0, err
	}

	return fodderAmount * fatteningFactor / float64(numberOfCows), nil
}

// TODO: define the 'ValidateInputAndDivideFood' function
func ValidateInputAndDivideFood(fodderCalculator FodderCalculator, numberOfCows int) (float64, error) {
	if numberOfCows > 0 {
		return DivideFood(fodderCalculator, numberOfCows)
	}

	return 0, ErrInvalidNumberOfCows
}

// TODO: define the 'ValidateNumberOfCows' function
func ValidateNumberOfCows(numberOfCows int) error {
	if numberOfCows < 0 {
		return &InvalidCowsError{
			numberOfCows: numberOfCows,
			message: "there are no negative cows",
		}
	}

	if numberOfCows == 0 {
		return &InvalidCowsError{
			numberOfCows: numberOfCows,
			message: "no cows don't need food",
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
