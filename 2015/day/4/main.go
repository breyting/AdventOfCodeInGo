package main

import (
	"crypto/md5"
	"fmt"
	"io"
)

func main() {
	//Part 1
	res := hashStartingWithFiveZero("yzbqklnj")
	fmt.Println(res)

	//Part 2
	res = hashStartingWithSixZero("yzbqklnj")
	fmt.Println(res)
}

func hashStartingWithFiveZero(input string) int {
	num := 0
	hash := "Starting string doesn't matter, just need to set up for the for loop"
	for hash[:5] != "00000" {
		h := md5.New()
		num += 1
		test := fmt.Sprintf("%s%d", input, num)
		io.WriteString(h, test)
		hash = fmt.Sprintf("%x", h.Sum(nil))
	}
	return num
}

func hashStartingWithSixZero(input string) int {
	num := 0
	hash := "Starting string doesn't matter, just need to set up for the for loop"
	for hash[:6] != "000000" {
		h := md5.New()
		num += 1
		test := fmt.Sprintf("%s%d", input, num)
		io.WriteString(h, test)
		hash = fmt.Sprintf("%x", h.Sum(nil))
	}
	return num
}
