package elon

import "fmt"

// Your first steps could be to read through the tasks, and create
// these functions with their correct parameter lists and return types.
// The function body only needs to contain `panic("")`.
//
// This will make the tests compile, but they will fail.
// You can then implement the function logic one by one and see
// an increasing number of tests passing as you implement more
// functionality.

// TODO: define the 'Drive()' method
func (car *Car) Drive() {
	if(car.battery < car.batteryDrain){
		return
	}

	car.distance += car.speed
	car.battery -= car.batteryDrain
}

// TODO: define the 'DisplayDistance() string' method
func (car Car) DisplayDistance() string {
	return fmt.Sprintf("Driven %d meters", car.distance)
}

func (car Car) DisplayBattery() string {
	return fmt.Sprintf("Battery at %d%%", car.battery)
}


// TODO: define the 'CanFinish(trackDistance int) bool' method
func (car Car) CanFinish(trackDistance int) bool {
	var timeToFinish = float64(trackDistance) / float64(car.speed)
	var batteryUsed  = timeToFinish * float64(car.batteryDrain)
	var remainingBattery = float64(car.battery) - batteryUsed

	return remainingBattery >= 0
}