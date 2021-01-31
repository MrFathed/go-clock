package main

import (
	"fmt"
	"time"

	"github.com/inancgumus/screen"
)

func main() {
	screen.Clear()

	for {
		screen.MoveTopLeft()

		now := time.Now()
		hour, minute, second := now.Hour(), now.Minute(), now.Second()

		clock := [...]placeholder{
			digits[hour/10], digits[hour%10],
			colon,
			digits[minute/10], digits[minute%10],
			colon,
			digits[second/10], digits[second%10],
		}

		for i := 0; i < 5; i++ {
			for _, digit := range clock {
				fmt.Print(digit[i], " ")
			}
			fmt.Println()
		}
		fmt.Println()

		time.Sleep(time.Second)
	}
}
