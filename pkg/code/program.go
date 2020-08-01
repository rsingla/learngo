package code

import (
	"fmt"
)

func main() {
	fmt.Println("Hello, This is Rajeev")
	fmt.Printf("%f\n", getMean())
	fizzBuzz(10)
	response := evenEnded(11)
	fmt.Println(response)

	maximalNumber()

	countWords()

	arr := []int{2, 3, 4, 5, 6}
	updateViaPointer(&arr)
}

func getMean() float64 {

	var x float64
	var y float64

	x = 13.50
	y = 12

	fmt.Printf("x %v value is and format of x %T\n", x, x)
	fmt.Printf("y %v value is and format of y %T\n", y, y)

	var mean float64

	mean = (x + y) / 2

	fmt.Printf("mean %v value is and format of mean %T\n", mean, mean)

	return mean
}

func fizzBuzz(start int) {

	var i int = 1

	for i = 1; i <= start; i++ {
		if i%5 == 0 && i%3 == 0 {
			fmt.Printf("fizzbuzz\n")
		} else if i%5 == 0 {
			fmt.Printf("buzz\n")
		} else if i%3 == 0 {
			fmt.Printf("fizz\n")
		} else {
			fmt.Printf("%d\n", i)
		}
	}

}

func evenEnded(num int) bool {

	var s int
	var e int

	s = num % 10
	num = num / 10

	for num > 0 {
		e = num % 10
		num = num / 10
	}

	if s == e {
		return true
	}

	return false
}

func maximalNumber() {
	var arr = []int{16, 12, 19, 29, 39, 10, 4}

	var max = arr[0]

	for _, s := range arr[1:] {
		if max < s {
			max = s
		}
	}

	fmt.Println(max)
}



func updateViaPointer(ptr *[]int) {
	var result *[]int
	for i,_ := len(*ptr) {
		*result[i] = *ptr[i] * 2
	}
}
