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

func mapBtwZeroAndSeven(num int) int {
	for num < 0 {
		num = 7 - num
	}

	return num % 7
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

	const (
		veryShortWeekShortness = 4 // very short are 2 charters long... add 2 extra characters
		shortWeekShortness     = 5 // short are 3 charters long... add 2 extra characters
		longWeekShortness      = 11 // long are at most 9 charters long... add 2 extra characters
	)
	weekShortness := shortWeekShortness
	weeks := weeksLong
	if weekShortness == veryShortWeekShortness {
		weeks = weeksVeryShort
	} else if weekShortness == shortWeekShortness {
		weeks = weeksShort
	}
	weeksLen := 7 * weekShortness - 2

	startsWithMonday := false
	dayOffset := 0
	if startsWithMonday {
		dayOffset = 1
	}

	currentTime := time.Now()
	timefmt := currentTime.Format("03:04:05 PM")
	monthYear := fmt.Sprintf("%s %d", currentTime.Month(), currentTime.Year())
	
	fmt.Println(strings.Repeat("-", weeksLen))
	fmt.Println()
	fmt.Printf("%*s\n", (weeksLen+len(monthYear))/2, monthYear)
	fmt.Printf("%*s\n", (weeksLen+len(monthYear))/2, timefmt)
	fmt.Println()
	fmt.Println(strings.Repeat("-", weeksLen))
	
	currentWeekDay := int(currentTime.Weekday())
	for i, _ := range weeks {
		idx := (i + dayOffset) % 7
		if idx == currentWeekDay {
			setRGBBackground(255, 255, 255)
			setRGBForeground(0, 0, 0)
		}
		fmt.Printf("%-*s", weekShortness-2 ,weeks[idx])
		resetStyleAndColors()
		fmt.Printf("  ")
	}
	fmt.Println()
	fmt.Println()

	firstOfMonth := time.Date(currentTime.Year(), currentTime.Month(), 1, 0, 0, 0, 0, currentTime.Location())
	lastOfMonth := firstOfMonth.AddDate(0, 1, -1)
	dayCount := lastOfMonth.Day()

	
	daysPassed := int(currentTime.AddDate(0, 0, -currentTime.Day() + 1).Weekday()) - dayOffset
	daysPassed = mapBtwZeroAndSeven(daysPassed)
	for range daysPassed {
		fmt.Printf("%-*s  ", weekShortness-2, "")
	}

	for i := range dayCount {
		if (daysPassed + i) % 7 == 0 {
			fmt.Println()
		}
		if i + 1 == currentTime.Day() {
			setRGBBackground(255, 255, 255)
			setRGBForeground(0, 0, 0)
		}
		fmt.Printf("%-*s", weekShortness-2, fmt.Sprintf("%d", i+1))
		resetStyleAndColors()
		fmt.Printf("  ")
	}

	fmt.Println()
	fmt.Println()
}
