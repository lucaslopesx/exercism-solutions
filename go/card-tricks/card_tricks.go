package cards

// FavoriteCards returns a slice with the cards 2, 6 and 9 in that order.
func FavoriteCards() []int {
	cards := []int {2, 6, 9}

	return cards;
}

// GetItem retrieves an item from a slice at given position.
// If the index is out of range, we want it to return -1.
func GetItem(slice []int, index int) int {
	if index < 0 {
		return -1;
	}

	length := len(slice)

	if(index >= length){
		return -1;
	}

	return slice[index]
}

// SetItem writes an item to a slice at given position overwriting an existing value.
// If the index is out of range the value needs to be appended.
func SetItem(slice []int, index, value int) []int {
	item := GetItem(slice, index)

	if(item != -1){
		slice[index] = value

		return slice
	}
	
	return append(slice, value)
}

// PrependItems adds an arbitrary number of values at the front of a slice.
func PrependItems(slice []int, values ...int) []int {
	values = append(values, slice...)

	return values
}

// RemoveItem removes an item from a slice by modifying the existing slice.
func RemoveItem(slice []int, index int) []int {
	item := GetItem(slice, index)

	if(item != -1){
		slice = append(slice[:index], slice[index+1:]...)

		return slice
	}

	return slice
}
