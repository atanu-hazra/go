package main

import (
	"fmt"
	"time" // Package time provides functionality for measuring and displaying time
)

func main() {
	fmt.Println("Time in golang") // output: Time in golang

	// 1. Get current local time
	presentTime := time.Now()
	fmt.Println(presentTime) // output: 2026-08-17 08:40:33.978978281 +0530 IST m=+0.000088922 (current date & time)

	// 2. Go's unique layout string:
	// Go uses a specific reference date for formatting instead of tokens like %Y-%m-%d:
	// Reference date: Mon Jan 2 15:04:05 MST 2006 (01/02 03:04:05PM '06 -0700)
	timeFormatter := "01-02-2006 Monday 15:04:05"

	// 3. Formatting present time with the layout string
	fmt.Println(presentTime.Format(timeFormatter)) // output: 08-17-2026 Monday 08:40:33

	// 4. Time arithmetic using .Add() and time.Duration constants
	nextHour := presentTime.Add(time.Hour).Format(timeFormatter)
	tomorrow := presentTime.Add(24 * time.Hour).Format(timeFormatter)

	fmt.Println(nextHour, tomorrow) // output: 08-17-2026 Monday 09:40:33 08-18-2026 Tuesday 08:40:33

	// 5. Unix timestamp: seconds passed since Unix epoch (Jan 1, 1970 UTC)
	fmt.Println(presentTime.Unix()) // output: 1786936233

	// 6. Memory address of variable 'tomorrow' using & (address-of operator)
	fmt.Println(&tomorrow) // output: 0xc0000bc010 (hex memory address)
}
