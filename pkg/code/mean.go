package code

import "fmt"

func getMean() float64 {

	var x float64
	var y float64

	x = 10.00
	y = 12.00

	fmt.Printf("x%v value is and format is x%T", x, x)
	fmt.Printf("y%v value is and format is y%T", y, y)

	var mean float64

	mean = x + y/2

	fmt.Printf("mean%v value is and format is mean%T", mean, mean)

	return mean
}
