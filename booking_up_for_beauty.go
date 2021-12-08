package booking

import (
	"fmt"
	"time"
)

// Schedule returns a time.Time from a string containing a date
func Schedule(date string) time.Time {
	//panic("Please implement the Schedule function")
	//Schedule("7/25/2019 13:45:00")
	//const myLayout = "2/29/2112 11:59:59"
	////myLayout := "7/25/2019 13:45:00"
	//t, _ := time.Parse(myLayout, date)
	const layout = "1/2/2006 15:04:05"
	//date := "12/8/2015 12:00:00"
	t, _ := time.Parse(layout, date)
	fmt.Println(t)
	return t
}

// HasPassed returns whether a date has passed
//isYear2000BeforeYear3000 := year2000.Before(year3000) // True
func HasPassed(date string) bool {
	now := time.Now()
	ourDate := Schedule(date)
	return ourDate.Before(now)
}

// IsAfternoonAppointment returns whether a time is in the afternoon
func IsAfternoonAppointment(date string) bool {
	//panic("Please implement the IsAfternoonAppointment function")
	//Thursday, July 25, 2019 13:45:00
	const layout = "Thursday, July 25, 2019 13:45:00"
	t, _ := time.Parse(layout, date)
	var hour, min, sec = t.Clock()
	if hour <= 11 {
		fmt.Print(hour, min, sec)
		return false
	}
	return true
}

// Description returns a formatted string of the appointment time
func Description(date string) string {
	//panic("Please implement the Description function")
	//"You have an appointment on Thursday, July 25, 2019, at 13:45."
	week := Schedule(date).Weekday()
	// get hours and minutes Clock() (hour, min, sec int)
	var hour, min, sec = Schedule(date).Clock()
	year, month, day := Schedule(date).Date()
	fmt.Print(sec)
	return fmt.Sprintf("You have an appointment on %s, %s %d, %d, at %d:%d.", week, month, day, year, hour, min)

}

// AnniversaryDate returns a Time with this year's anniversary
func AnniversaryDate() time.Time {
	//panic("Please implement the AnniversaryDate function")
	//2020-09-15
	//t = time.Date(2021, time.Month(2), 21, 1, 10, 30, 0, time.UTC)
	//fmt.Println(t)
	//
	////time.Parse() Example
	////Parse YYYY-MM-DD
	//t, _ = time.Parse("2006-01-02", "2020-01-29")
	//t :=time.Date(2020, time.Month(9), 15, 1, 10, 30, 0, time.UTC)
	t, _ := time.Parse("2006-01-02", "2021-09-15")
	return t
}
