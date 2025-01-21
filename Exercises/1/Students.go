package main

import (
	"GoLang/Exercises/1/students"
	"fmt"
	"time"
)

func main() {
	sMap := map[int]student.Student{
		1: {Name: "Mike Hunt", ID: 1, Age: 24, DOB: time.Date(1999, 3, 1, 0, 0, 0, 0, time.UTC)},
		2: {Name: "Anita Max Wynn", ID: 2, Age: 23, DOB: time.Date(2001, 3, 9, 0, 0, 0, 0, time.UTC)},
		3: {Name: "Ben Dover", ID: 3, Age: 48, DOB: time.Date(1976, 8, 1, 0, 0, 0, 0, time.UTC)},
	}

	for k, v := range sMap {
		fmt.Println("ID:", k, "-", v.String())
	}
}
