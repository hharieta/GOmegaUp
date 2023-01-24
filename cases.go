package main

import (
	"fmt"
	"time"
)

func main() {
	time_now := time.Now()

	switch today := time_now.Weekday(); today {
	case 0:
		fmt.Println("Descanso Domingo")
	case 1:
		fmt.Println("Odio los Lunes")
	default:
		fmt.Println("Trabajo")
	}
}
