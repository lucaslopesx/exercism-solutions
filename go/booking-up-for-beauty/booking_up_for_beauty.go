package booking

import (
	"fmt"
	"time"
)

// Schedule returns a time.Time from a string containing a date.
func Schedule(date string) time.Time {
	layout := "1/02/2006 15:04:05"
	
	t, err := time.Parse(layout, date)

	if(err != nil){
		panic(err)
	}

	return t
}

// HasPassed returns whether a date has passed.
func HasPassed(date string) bool {
	layout := "January 2, 2006 15:04:05"

	t, err := time.Parse(layout, date)

	if(err != nil){
		panic(err)
	}

	if(time.Now().Compare(t) == 1){
		return true
	}
	
	return false
}

// IsAfternoonAppointment returns whether a time is in the afternoon.
func IsAfternoonAppointment(date string) bool {
	layout := "Monday, January 2, 2006 15:04:05"

	t, err := time.Parse(layout, date)

	if(err != nil){
		panic(err)
	}

	if(t.Hour() >= 12 && t.Hour() < 18){
		return true
	}

	return false
}

// Description returns a formatted string of the appointment time.
func Description(date string) string {
	layout := "1/2/2006 15:04:05"
	format := "You have an appointment on Monday, January 2, 2006, at 15:04."

	t, err := time.Parse(layout, date)

	if(err != nil){
		panic(err)
	}

	return t.Format(format)
}

// AnniversaryDate returns a Time with this year's anniversary.
func AnniversaryDate() time.Time {
	layout := "2006-01-02"

	date := fmt.Sprintf("%d-09-15", time.Now().Year())

	t, err := time.Parse(layout, date)

	if(err != nil){
		panic(err)
	}

	return t
}
