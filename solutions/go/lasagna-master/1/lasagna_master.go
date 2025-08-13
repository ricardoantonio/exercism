package lasagna

// TODO: define the 'PreparationTime()' function
func PreparationTime(layers []string, timePerLayer int) int {
	if timePerLayer == 0 {
		timePerLayer = 2 // default time per layer
	}
	return len(layers) * timePerLayer
}

// TODO: define the 'Quantities()' function
func Quantities(layers []string) (int, float64) {
	noodles := 50
	sause := 0.2

	totalNoodles := 0
	totalSause := 0.0

	for i := 0; i < len(layers); i++ {
		switch layers[i] {
		case "noodles":
			totalNoodles += noodles
		case "sauce":
			totalSause += sause
		}
	}

	return totalNoodles, totalSause
}

// TODO: define the 'AddSecretIngredient()' function
func AddSecretIngredient(friendsList []string, myList []string) {
	if len(friendsList) == 0 || len(myList) == 0 {
		return
	}
	myList[len(myList)-1] = friendsList[len(friendsList)-1]
}

// TODO: define the 'ScaleRecipe()' function
func ScaleRecipe(quantities []float64, scale int) []float64 {
	scaledQuantities := make([]float64, len(quantities))

	for i := 0; i < len(quantities); i++ {
		scaledQuantities[i] = quantities[i] / 2.0 * float64(scale)
	}

	return scaledQuantities
}

// Your first steps could be to read through the tasks, and create
// these functions with their correct parameter lists and return types.
// The function body only needs to contain `panic("")`.
//
// This will make the tests compile, but they will fail.
// You can then implement the function logic one by one and see
// an increasing number of tests passing as you implement more
// functionality.
