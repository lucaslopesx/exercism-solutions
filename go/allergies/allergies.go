package allergies

import "math"

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

func GetClosestAllergyValue(allergies uint) uint {
	var closestAllergyValue uint
	i := 0
	for {
		closestAllergyValue = uint(math.Pow(2, float64(i)))
		if(closestAllergyValue > uint(allergies)){
			closestAllergyValue = uint(math.Pow(2, float64(i - 1)))
			break
		}
		i++
	}
	return closestAllergyValue
}


func Allergies(allergies uint) []string {
	if allergies == 0 {
		return []string{} 
	}
	
	v, exists := allergiesMap[allergies]
	if exists {
		return []string{v}
	}

	var res []string

	closestAllergyValue := GetClosestAllergyValue(allergies)

	sub := allergies - closestAllergyValue

	if(sub == 0){
		closestAllergy := allergiesMap[closestAllergyValue]
		return []string{closestAllergy}
	}

	closestAllergy2, exists4 := allergiesMap[closestAllergyValue]
		
	if(exists4){
		res = append(res, closestAllergy2)
	}
	
	for {
		newClosestAllergyValue := GetClosestAllergyValue(sub)
		newClosestAllergy, exists3 := allergiesMap[newClosestAllergyValue]
		if(exists3){
			res = append(res, newClosestAllergy)
		}
		sub = sub - newClosestAllergyValue
		if sub <= 0 {
			break
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
