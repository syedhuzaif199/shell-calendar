package main

import (
	"fmt"
	"strings"
	"time"
)

type EscCode string

func getRGBForegroundEscCode(r, g, b uint8) EscCode {
	return EscCode(fmt.Sprintf("\x1B[38;2;%d;%d;%dm", r, g, b))
}

func getRGBBackgroundEscCode(r, g, b uint8) EscCode {
	return EscCode(fmt.Sprintf("\x1B[48;2;%d;%d;%dm", r, g, b))
}

func setRGBForeground(r, g, b uint8) {
	fmt.Printf("\x1B[38;2;%d;%d;%dm", r, g, b)
}

func setRGBBackground(r, g, b uint8) {
	fmt.Printf("\x1B[48;2;%d;%d;%dm", r, g, b)
}

func resetStyleAndColors() {
	fmt.Printf("\x1B[0m")
}

func main() {
	weeksVeryShort := []string{
		"Su",
		"Mo",
		"Tu",
		"We",
		"Th",
		"Fr",
		"Sa",
	}
	weeksShort := []string{
		"Sun",
		"Mon",
		"Tue",
		"Wed",
		"Thu",
		"Fri",
		"Sat",
	}
	weeksLong := []string{
		"Sunday",
		"Monday",
		"Tuesday",
		"Wednesday",
		"Thursday",
		"Friday",
		"Saturday",
	}

	startsWithMonday := false
	const (
		veryShortWeekShortness = 4 // long are at most 7 charters long... add 2 extra characters
		shortWeekShortness     = 5 // very short are 2 charters long... add 2 extra characters
		longWeekShortness      = 9 // short are 3 charters long... add 2 extra characters
	)
	weekShortness := shortWeekShortness
	weeks := weeksLong
	if weekShortness == veryShortWeekShortness {
		weeks = weeksVeryShort
	} else if weekShortness == shortWeekShortness {
		weeks = weeksShort
	}
	weeksLen := 0
	for _, week := range weeks {
		weeksLen += len(week)
		weeksLen += 2
	}
	weeksLen -= 2

	dayOffset := 0
	if startsWithMonday {
		dayOffset = 1
	}

	currentTime := time.Now()
	// timefmt := currentTime.Format("12:35:25 AM");
	// format : 12 hour
	timefmt := currentTime.Format("03:04:05 PM")
	monthYear := fmt.Sprintf("%s %d", currentTime.Month(), currentTime.Year())
	formatStr := fmt.Sprintf("%%%ds", (weeksLen+len(monthYear))/2)
	currentWeekDay := int(currentTime.Weekday())

	fmt.Println(strings.Repeat("-", weeksLen))
	fmt.Println()
	fmt.Printf(formatStr+"\n", monthYear)
	fmt.Printf(formatStr+"\n", timefmt)
	fmt.Println()
	fmt.Println(strings.Repeat("-", weeksLen))
	formatStr = fmt.Sprintf("%%-%ds", weekShortness-2)
	for i, _ := range weeks {
		idx := (i + dayOffset) % 7
		if idx == currentWeekDay {
			setRGBBackground(255, 255, 255)
			setRGBForeground(0, 0, 0)
		}
		fmt.Printf(formatStr, weeks[idx])
		resetStyleAndColors()
		fmt.Printf("  ")
	}
	fmt.Println()
	dayCount := 30 + (int(currentTime.Month())+1)%2
	if currentTime.Month() == time.February {
		dayCount -= 3
		if currentTime.Year()%4 == 0 {
			dayCount += 1
		}
	}

	
	for i := 0; i < currentWeekDay; i++ {
		fmt.Printf(formatStr, "")
		fmt.Printf("  ")
	}

	for i := 0; i < dayCount; i++ {
		if (currentWeekDay + i) % 7 == 0 {
			fmt.Println()
		}
		if i + 1 == currentTime.Day() {
			setRGBBackground(255, 255, 255)
			setRGBForeground(0, 0, 0)
		}
		fmt.Printf(formatStr, fmt.Sprintf("%d", i+1))
		resetStyleAndColors()
		fmt.Printf("  ")
	}


}
