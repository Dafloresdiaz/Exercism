// Package weather to know the current weather condition.
package weather

//CurrentCondition to know the current  condition.
var CurrentCondition string

//CurrentLocation to know the current location.
var CurrentLocation string

//Forecast return the current location and the current condition.
func Forecast(city, condition string) string {
	CurrentLocation, CurrentCondition = city, condition
	return CurrentLocation + " - current weather condition: " + CurrentCondition
}
