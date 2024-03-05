package lasagna

// TODO: define the 'PreparationTime()' function
func PreparationTime(layers []string, avgPrepTime int) int {
	time := 0

	if(avgPrepTime == 0){
		avgPrepTime = 2
	}

	for i := 0; i < len(layers); i++ {
		time += avgPrepTime
	}

	return time
}

// TODO: define the 'Quantities()' function
func Quantities(layers []string) (int, float64) {
	noodle := 0
	sauce := 0.0

	for i := 0; i < len(layers); i++ {
		if(layers[i] == "noodles"){
			noodle += 50
		}

		if(layers[i] == "sauce"){
			sauce += 0.2
		}
	}

	return noodle, sauce
}

// TODO: define the 'AddSecretIngredient()' function
func AddSecretIngredient(friendsList []string, myList []string) {
	friendIngredient := friendsList[len(friendsList) - 1]

	myList[len(myList) - 1] = friendIngredient
}

// TODO: define the 'ScaleRecipe()' function
func ScaleRecipe(quantities []float64, portions int) []float64 {
	result := append([]float64{}, quantities...)

	for i := 0; i < len(result); i++ {
		onePortion := result[i] / 2
		result[i] = onePortion * float64(portions)
	}

	return result
}

// Your first steps could be to read through the tasks, and create
// these functions with their correct parameter lists and return types.
// The function body only needs to contain `panic("")`.
// 
// This will make the tests compile, but they will fail.
// You can then implement the function logic one by one and see
// an increasing number of tests passing as you implement more 
// functionality.
