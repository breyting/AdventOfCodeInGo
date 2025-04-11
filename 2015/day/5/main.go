package main

import (
	"fmt"
	"slices"
	"strings"

	config "github.com/breyting/AdventOfCodeInGo"
)

func main() {
	// Part 1
	input, err := config.ReadInputAOC("2015", "5")
	if err != nil {
		panic("FUCK")
	}
	res := NumberOfNiceString(input)
	fmt.Println(res)

	// Part 2
	res = NumberOfNewNiceString(input)
	fmt.Println(res)
}

func NumberOfNiceString(input string) int {
	listOfStrings := strings.Split(input, "\n")
	total := 0
	vowel := []rune{'a', 'e', 'i', 'o', 'u'}
	badString := []string{"ab", "cd", "pq", "xy"}
	for _, str := range listOfStrings {
		if len(str) == 1 {
			break
		}

		test := isNiceString{
			numberOfVowel:                 0,
			containsSameLetterTwiceInARow: false,
			notContainsBadString:          true,
		}

		for i, char := range str {
			if slices.Contains(vowel, char) {
				test.numberOfVowel += 1
			}
			if i != len(str)-1 {
				if char == rune(str[i+1]) {
					test.containsSameLetterTwiceInARow = true
				}
				bothLetters := fmt.Sprintf("%s%s", string(char), string(str[i+1]))
				if slices.Contains(badString, bothLetters) {
					test.notContainsBadString = false
				}
			}
		}

		if test.numberOfVowel >= 3 &&
			test.containsSameLetterTwiceInARow &&
			test.notContainsBadString {
			total += 1
		}
	}

	return total
}

type isNiceString struct {
	numberOfVowel                 int
	containsSameLetterTwiceInARow bool
	notContainsBadString          bool
}

func NumberOfNewNiceString(input string) int {
	listOfStrings := strings.Split(input, "\n")
	total := 0

	for _, str := range listOfStrings {
		doubleLetter := []string{}

		twicePairOfLetter := false
		letterRepeat := false

		for i, char := range str {
			if i < len(str)-1 && !twicePairOfLetter {
				bothLetters := fmt.Sprintf("%s%s", string(char), string(str[i+1]))
				if slices.Contains(doubleLetter, bothLetters) && slices.Index(doubleLetter, bothLetters) != i-1 {
					twicePairOfLetter = true
				} else {
					doubleLetter = append(doubleLetter, bothLetters)
				}
			}

			if i < len(str)-2 && !letterRepeat && char == rune(str[i+2]) {
				letterRepeat = true
			}

			if twicePairOfLetter && letterRepeat {
				total += 1
				break
			}
		}
	}

	return total
}
