package birdwatcher

// TotalBirdCount return the total bird count by summing
// the individual day's counts.
func TotalBirdCount(birdsPerDay []int) int {
	var sum int

	for i := 0; i < len(birdsPerDay); i++ {
		sum += birdsPerDay[i]
	}

	return sum
}

// BirdsInWeek returns the total bird count by summing
// only the items belonging to the given week.
func BirdsInWeek(birdsPerDay []int, week int) int {
	lastDayOfWeek := week * 7
	firstDayOfWeek := lastDayOfWeek - 7

	birdsPerDay = birdsPerDay[firstDayOfWeek:lastDayOfWeek]

	return TotalBirdCount(birdsPerDay)
}

// FixBirdCountLog returns the bird counts after correcting
// the bird counts for alternate days.
func FixBirdCountLog(birdsPerDay []int) []int {
	flag := true
	for i := 0; i < len(birdsPerDay); i++ {
		if(flag){
			birdsPerDay[i] += 1
		}
		flag = !flag
	}

	return birdsPerDay
}
