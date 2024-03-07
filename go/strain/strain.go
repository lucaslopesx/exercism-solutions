package strain

// Implement the "Keep" and "Discard" function in this file.
func Keep[T any](list []T, filterFunc func(T) bool) []T {
	var items []T;

	for i := 0; i < len(list); i++ {
		if filterFunc(list[i]) {
			items = append(items, list[i])
		}
	}
	
	return items
} 

func Discard[T any](list []T, filterFunc func(T) bool) []T {
	var items []T;

	for i := 0; i < len(list); i++ {
		if !filterFunc(list[i]) {
			items = append(items, list[i])
		}
	}
	
	return items
}

// You will need typed parameters (aka "Generics") to solve this exercise.
// They are not part of the Exercism syllabus yet but you can learn about
// them here: https://go.dev/tour/generics/1
