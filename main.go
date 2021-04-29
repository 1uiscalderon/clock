package main

import (
	"fmt"
	"time"
)

func main() {
	type number [5]string

	zero := number{
		"███",
		"█ █",
		"█ █",
		"█ █",
		"███",
	}

	one := number{
		"██ ",
		" █ ",
		" █ ",
		" █ ",
		"███",
	}

	two := number{
		"███",
		"  █",
		"███",
		"█  ",
		"███",
	}

	three := number{
		"███",
		"  █",
		"███",
		"  █",
		"███",
	}

	four := number{
		"█ █",
		"█ █",
		"███",
		"  █",
		"  █",
	}

	five := number{
		"███",
		"█  ",
		"███",
		"  █",
		"███",
	}

	six := number{
		"███",
		"█  ",
		"███",
		"█ █",
		"███",
	}

	seven := number{
		"███",
		"  █",
		"  █",
		"  █",
		"  █",
	}

	eight := number{
		"███",
		"█ █",
		"███",
		"█ █",
		"███",
	}

	nine := number{
		"███",
		"█ █",
		"███",
		"  █",
		"███",
	}

	colon := number{
		"   ",
		" ░ ",
		"   ",
		" ░ ",
		"   ",
	}

	digits := [10]number{
		zero, one, two, three, four, five, six, seven, eight, nine,
	}
	fmt.Print("\033[2J")

	for {
		fmt.Print("\033[H")

		now := time.Now()
		hour, min, sec := now.Hour(), now.Minute(), now.Second()

		clock := []number{
			digits[hour/10], digits[hour%10],
			colon,
			digits[min/10], digits[min%10],
			colon,
			digits[sec/10], digits[sec%10],
		}

		for line := 0; line < 5; line++ {
			for digit := 0; digit < 8; digit++ {
				fmt.Print(clock[digit][line], "  ")
			}
			fmt.Println()
		}

		// pause for 1 second
		time.Sleep(time.Second)
	}
}
