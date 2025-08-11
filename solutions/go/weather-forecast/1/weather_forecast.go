// Package weather forecast the current weather condition of various cities in Goblinocus.
package weather

// CurrentCondition is the current weather condition in the city.
var CurrentCondition string

// CurrentLocation is the name of the city for which the weather is being forecasted.
var CurrentLocation string

// Forecast returns the current weather condition for a given city.
func Forecast(city, condition string) string {
	CurrentLocation, CurrentCondition = city, condition
	return CurrentLocation + " - current weather condition: " + CurrentCondition
}
