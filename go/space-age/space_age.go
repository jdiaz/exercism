package space

import (
	"fmt"
	"strings"
)

type Planet string

const earthObitTime float64 = 31557600.00

func Age(seconds float64, p Planet) float64 {
	switch strings.ToLower(string(p)) {
	case "earth":
		return seconds / earthObitTime
	case "mercury":
		return seconds / (earthObitTime * 0.2408467)
	case "venus":
		return seconds / (earthObitTime * 0.61519726)
	case "mars":
		return seconds / (earthObitTime * 1.8808158)
	case "jupiter":
		return seconds / (earthObitTime * 11.862615)
	case "saturn":
		return seconds / (earthObitTime * 29.447498)
	case "uranus":
		return seconds / (earthObitTime * 84.016846)
	case "neptune":
		return seconds / (earthObitTime * 164.79132)
	default:
		fmt.Println("No love for Pluto")
	}

	return 0.0
}
