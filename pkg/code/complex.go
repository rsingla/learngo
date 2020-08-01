package code


func AddComplex() complex64 {
	var n1 complex64 = 1 + 2i
	var n2 complex64 = 2 + 4i

	result := n1 + n2

	return result
}

func SumComplex(c1 complex64, c2 complex64) complex64 {
	result := c1 + c2
	return result
}

func SubtractComplex( c1 complex64, c2 complex64) complex64 {
	result := c1 - c2
	return result
}