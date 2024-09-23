package main

import (
	"fmt"
	"strings"
)

func task_1() {
	numbers := [5]int{2, 4, 6, 8, 10}
	for i := 0; i < 5; i++ {
		fmt.Println("Квадрат числа", numbers[i], ":", numbers[i]*numbers[i])
	}
}

func task_2() {
	temperatures := []float64{-25.4, -27.0, 13.0, 19.0, 15.5, 24.5, -21.0, 32.5}
	groups := make(map[int][]float64)
	for _, temp := range temperatures {
		groupKey := int(temp/10) * 10
		groups[groupKey] = append(groups[groupKey], temp)
	}
	fmt.Println(groups)
}

func task_3() string {

	s := "abcd"
	var reversed []rune

	for _, r := range s {
		reversed = append([]rune{r}, reversed...)
	}

	return string(reversed)
}
func task_4() {
	word := "hot pot dont"
	words := strings.Fields(word)
	for i, j := 0, len(words)-1; i < j; i, j = i+1, j-1 {
		words[i], words[j] = words[j], words[i]
		words := strings.Join(words, " ")
		fmt.Println(words)
	}
}

func task_5(s string) bool {
	s = strings.ToLower(s)
	charMap := make(map[rune]bool)

	for _, char := range s {
		if char != ' ' {
			if charMap[char] {
				return false
			}
			charMap[char] = true
		}
	}
	return true
}

func main() {
	// task_1()
	// task_2()
	// task_3()
	// task_4()
	// fmt.Println(task_5("asdf"))
	// fmt.Println(task_5("aaws"))

}
