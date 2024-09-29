package main

import (
	"fmt"
	"time"
)

func main() {
	i := 1
	fmt.Print("Write ", i, " as ")
	switch i {
	case 1:
		fmt.Println("ONE")
	case 2:
		fmt.Println("TWO")
	case 3:
		fmt.Println("THREE")
	default:
		fmt.Println("???")
	}

	switch time.Now().Weekday() {
	case time.Saturday, time.Sunday:
		fmt.Println("It's the weekend")
	default:
		fmt.Println("It's the weekday")
	}

	whatAmI := func(i interface{}) {
		switch t := i.(type) {
		case bool:
			fmt.Println("BOOLEAN")
		case int:
			fmt.Println("INT")
		default:
			fmt.Printf("Don't know type %T\n", t)
		}
	}

	whatAmI(true)
	whatAmI(1)
	whatAmI("string or what")

	whatAmI2(true)
	whatAmI2(1)
	whatAmI2("string or what")
}

func whatAmI2(i interface{}) {
	switch t := i.(type) {
	case bool:
		fmt.Println("BOOLEAN")
	case int:
		fmt.Println("INT")
	default:
		fmt.Printf("Don't know type %T\n", t)
	}
}
