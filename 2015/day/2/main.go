package main

import (
	"fmt"
	"strconv"
	"strings"

	config "github.com/breyting/AdventOfCodeInGo"
)

func main() {
	// Part 1
	input, _ := config.ReadInputAOC("2015", "2")
	res := totalWrappingPaper(input)
	fmt.Println(res)

	// Part 2
	res = totalFeetOfRibon(input)
	fmt.Println(res)
}

func totalWrappingPaper(input string) int {
	listOfGifts := strings.Split(input, "\n")
	sum := 0
	for _, gift := range listOfGifts {
		splitGift := strings.Split(gift, "x")
		// when hit empty list
		if len(splitGift) == 1 {
			continue
		}
		l, _ := strconv.Atoi(splitGift[0])
		w, _ := strconv.Atoi(splitGift[1])
		h, _ := strconv.Atoi(splitGift[2])

		arealw := 2 * l * w
		areawh := 2 * w * h
		areahl := 2 * h * l
		minArea := findMin(arealw, areawh, areahl)

		subTotal := arealw + areawh + areahl + (minArea / 2)
		sum += subTotal
	}
	return sum
}

func findMin(inputs ...int) int {
	min := inputs[0]
	for _, input := range inputs[1:] {
		if input < min {
			min = input
		}
	}
	return min
}

func totalFeetOfRibon(input string) int {
	listOfGifts := strings.Split(input, "\n")
	sum := 0
	for _, gift := range listOfGifts {
		splitGift := strings.Split(gift, "x")
		// when hit empty list
		if len(splitGift) == 1 {
			break
		}
		l, _ := strconv.Atoi(splitGift[0])
		w, _ := strconv.Atoi(splitGift[1])
		h, _ := strconv.Atoi(splitGift[2])

		max := findMax(l, w, h)

		feetToWrap := (2 * l) + (2 * w) + (2 * h) - (2 * max)
		feetForBow := l * w * h
		sum += feetToWrap + feetForBow
	}
	return sum
}

func findMax(inputs ...int) int {
	max := inputs[0]
	for _, input := range inputs[1:] {
		if input > max {
			max = input
		}
	}
	return max
}
