package allergies

var allergiesMap = map[uint]string {
	1: "eggs",
	2: "peanuts",
	4: "shellfish",
	8: "strawberries",
	16: "tomatoes",
	32: "chocolate",
	64: "pollen",
	128: "cats",
}

func Allergies(allergies uint) []string {
	var res []string
	for allergyValue, allergyName := range allergiesMap {
		if allergies&allergyValue != 0 {
			res = append(res, allergyName)
		}	
	}
	return res
}

func AllergicTo(allergies uint, allergen string) bool {
	allAllergies := Allergies(allergies)

	for _, v := range allAllergies {
		if(v == allergen){
			return true
		}
	}
	
	return false
}
