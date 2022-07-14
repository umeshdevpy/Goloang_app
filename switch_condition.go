package main

import (
	"fmt"
	"time"
)

func main(){
	fmt.Printf("How may days are remaining for Saturday!!\n")
	todayDay := time.Now().Weekday()

	switch time.Saturday {
	case todayDay + 0:
		fmt.Printf("Today..\n")
	case todayDay + 1:
		fmt.Printf("Tomorrow..\n")
	case todayDay + 2:
		fmt.Printf("In Two days..\n")
	default:
		fmt.Printf("It's long.")
	}
}

