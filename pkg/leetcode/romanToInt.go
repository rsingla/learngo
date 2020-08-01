package leetcode

func romanToInt(s string) int {
	
	var value int
	lastval := 0

	var roman = map[string]int{"I": 1,
		"V": 5,
		"X": 10,
		"L": 50,
		"C": 100,
		"D": 500,
		"M": 1000,
	}

	for _, v := range s {
		myVal := string(v)
		if lastval < roman[myVal] {
			value = value + (roman[myVal] - lastval) - lastval
		} else {
			value = value + roman[myVal]
		}
		lastval = roman[myVal]
	}
	return value
}
