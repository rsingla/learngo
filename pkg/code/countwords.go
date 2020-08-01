package code

import "fmt"

func countWords() {
	text := "Needs and pins Needs and pins sew me a catch sew me a sail to catch me a the wind"

	var words = make(map[string]int)
	var s string
	for _, wd := range text {

		if wd != ' ' {
			val := string(wd)
			s = s + val
		} else {
			words[s] = words[s] + 1
			s = ""
		}
	}

	fmt.Println(words)

	for key, word := range words {
		fmt.Println(key, " ", word)
	}

}
