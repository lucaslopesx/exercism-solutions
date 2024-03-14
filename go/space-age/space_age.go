package space

type Planet string

var earthYearsMap = map[Planet] float64 {
	"mercury": 0.2408467,
	"venus" : 0.61519726,
	"earth" : 1,
	"mars" : 1.88081580,
	"jupiter" : 11.8626150,
	"saturn" : 29.4474980,
	"uranus" : 84.0168460,
	"neptune" : 164.791320,
}

func Age(seconds float64, planet Planet) float64 {
	years, exists := earthYearsMap[planet]

	if !exists {
		return -1
	}

	planetSeconds := years * 31557600

	return seconds / planetSeconds	
}
