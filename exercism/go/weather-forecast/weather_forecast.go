// Package weather provides weather forecast information.
package weather

// CurrentCondition stores the current weather condition.
var CurrentCondition string
// CurrentLocation stores the current location to be forecasted.
var CurrentLocation string

// Forecast provides the weather forecast of a certain location.
func Forecast(city, condition string) string {
	CurrentLocation, CurrentCondition = city, condition
	return CurrentLocation + " - current weather condition: " + CurrentCondition
}
