package code

import (
	"fmt"
	"mathprog"
	"strconv"
)

var i int = 27

func main() {
	fmt.Printf("%v \n", i)
	var i int = 47
	fmt.Printf("%v \n", i)
	movie()
	convert()
	fmt.Println("--------")
	mathfunc()
	fmt.Println(mathprog.AddComplex())
	fmt.Println(mathprog.SumComplex(1+2i, 2+3i))
	fmt.Println(mathprog.SubtractComplex(4+4i, 2+3i))
}

var (
	actorName string = "Rozer stone"
	companion string = "Sarah Jane Smith"
	doctorNum int    = 3
	season    int    = 11
)

func movie() {
	fmt.Printf("%v, %T \n", actorName, actorName)
	fmt.Printf("%v, %T \n", companion, companion)
	fmt.Printf("%v, %T \n", doctorNum, doctorNum)
	fmt.Printf("%v, %T \n", season, season)

}

func mathfunc() {
	a := 10 // 1010
	b := 3  // 0011
	fmt.Println(a + b)
	fmt.Println(a - b)
	fmt.Println(a * b)
	fmt.Println(float32(a) / float32(b))
	fmt.Println(a % b)
	fmt.Println("--------")
	fmt.Println(a & b)
	fmt.Println(a | b)
	fmt.Println(a ^ b)
	fmt.Println(a &^ b)

}

func convert() {
	//Lower to Upper
	var sum int = 42
	fmt.Printf("%v, %T \n", sum, sum)
	var sumfloat float64 = float64(sum)
	fmt.Printf("%v, %T \n", sumfloat, sumfloat)

	//Upper to Lower
	var minus float64 = 22.90
	fmt.Printf("%v, %T \n", minus, minus)
	var minusint int = int(minus)
	fmt.Printf("%v, %T \n", minusint, minusint)

	v := "3.1415926535"
	if s, err := strconv.ParseFloat(v, 32); err == nil {
		fmt.Printf("%T, %v\n", s, s)
	}

	//Upper to Lower
	var myval float64 = 22.90
	fmt.Printf("%v, %T \n", myval, myval)
	myvalstr := strconv.FormatFloat(myval, 'g', -1, 64)

	fmt.Printf("%v, %T \n", myvalstr, myvalstr)

}
