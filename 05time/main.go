package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Time in golang")

	presentTime := time.Now()
	fmt.Println((presentTime))

	timeFormatter := "01-02-2006 Monday 15:04:05"
	
	fmt.Println(presentTime.Format(timeFormatter))
	nextHour := presentTime.Add(time.Hour).Format(timeFormatter)
	tomorrow := presentTime.Add(24 * time.Hour).Format(timeFormatter)

	fmt.Println(nextHour, tomorrow)
	fmt.Println(presentTime.Unix())
}
