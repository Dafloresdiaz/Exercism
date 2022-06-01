package lasagna

// TODO: define the 'OvenTime' constant
const OvenTime int = 40

// RemainingOvenTime returns the remaining minutes based on the `actual` minutes already in the oven.
func RemainingOvenTime(actualMinutesInOven int) int {
	var calculate int = OvenTime - actualMinutesInOven
	return calculate
}

// PreparationTime calculates the time needed to prepare the lasagna based on the amount of layers.
func PreparationTime(numberOfLayers int) int {
	var calculate int = 2 * numberOfLayers
	return calculate
}

// ElapsedTime calculates the total time needed to create and bake a lasagna.
func ElapsedTime(numberOfLayers, actualMinutesInOven int) int {
	var timeLayers int = PreparationTime(numberOfLayers)
	var calculate int = timeLayers + actualMinutesInOven
	return calculate
}
