package main

import (
	"fmt"
	"time"
)

func printSwitch() {
	switch time.Now().Weekday() {
	case time.Saturday, time.Sunday:
		fmt.Print("WEEKEND")
	default:
		fmt.Print(("WEEKDAY"))
	}
}
