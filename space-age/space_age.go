package space

// Planet is type
type Planet string

// Age document goes here
func Age(ageEarth float64, planet Planet) float64 {
	// return math.Round(((ageEarth / 31557600) * 100 / orbitalPeriod(planet))) / 100
	return ageEarth / (31557600 * orbitalPeriod(planet))
}

func orbitalPeriod(planet Planet) float64 {
	switch planet {
	case "Mercury":
		return 0.2408467
	case "Venus":
		return 0.6151972
	case "Mars":
		return 1.8808158
	case "Jupiter":
		return 11.862615
	case "Saturn":
		return 29.447498
	case "Uranus":
		return 84.016846
	case "Neptune":
		return 164.79132
	default:
		return 1.0
	}
}
