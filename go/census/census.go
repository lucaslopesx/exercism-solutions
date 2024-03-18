// Package census simulates a system used to collect census data.
package census

// Resident represents a resident in this city.
type Resident struct {
	Name    string
	Age     int
	Address map[string]string
}

// NewResident registers a new resident in this city.
func NewResident(name string, age int, address map[string]string) *Resident {
	return &Resident{
		Name: name,
		Age: age,
		Address: address,
	}
}

// HasRequiredInfo determines if a given resident has all of the required information.
func (r *Resident) HasRequiredInfo() bool {
	if r.Address == nil || r.Name == "" {
		return false
	}

	if len(r.Address) == 0 {
		return false
	}

	value, exists := r.Address["street"]

	if value == "" {
		return false
	}
	
	return exists
}

// Delete deletes a resident's information.
func (r *Resident) Delete() {
	*r = Resident{}
}

// Count counts all residents that have provided the required information.
func Count(residents []*Resident) int {
	count := 0
	for _, r := range residents {
		if r != nil && r.HasRequiredInfo() {
			count++
		}
	}
	return count
}
