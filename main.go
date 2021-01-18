package main

import (
	"fmt"
	"time"
)

func main() {
	now := time.Now()

	hour := now.Hour()
	minute := now.Minute()
	second := now.Second()
	period := "AM"

	if hour > 12 {
		hour = hour - 12
		period = "PM"
	}

	month := now.Month()
	day := now.Day()
	year := now.Year()

	fmt.Printf("%d/%d/%d\t%d:%02d:%02d %s\n",
		month,
		day,
		year,
		hour,
		minute,
		second,
		period,
	)
}
