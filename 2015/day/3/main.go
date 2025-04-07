package main

import (
	"fmt"
	"slices"

	config "github.com/breyting/AdventOfCodeInGo"
)

func main() {
	// Part 1
	input, _ := config.ReadInputAOC("2015", "3")
	res := numberOfHouseVisited(input)
	fmt.Println(res)

	// Part 2
	res = numberOfHouseVisitedWithRobot(input)
	fmt.Println(res)
}

func numberOfHouseVisited(input string) int {
	visited := make([]coord, 0)
	coord := coord{x: 0, y: 0}
	visited = append(visited, coord)
	for _, char := range input {
		switch char {
		case '^':
			coord.y += 1
			if !slices.Contains(visited, coord) {
				visited = append(visited, coord)
			}
		case 'v':
			coord.y -= 1
			if !slices.Contains(visited, coord) {
				visited = append(visited, coord)
			}
		case '>':
			coord.x += 1
			if !slices.Contains(visited, coord) {
				visited = append(visited, coord)
			}
		case '<':
			coord.x -= 1
			if !slices.Contains(visited, coord) {
				visited = append(visited, coord)
			}
		}
	}
	return len(visited)
}

type coord struct {
	x int
	y int
}

func numberOfHouseVisitedWithRobot(input string) int {
	visited := make([]coord, 0)
	coordSanta := coord{x: 0, y: 0}
	coordRobot := coord{x: 0, y: 0}
	visited = append(visited, coordSanta)
	for i, char := range input {
		switch char {
		case '^':
			if i%2 == 0 {
				coordRobot.y += 1
				if !slices.Contains(visited, coordRobot) {
					visited = append(visited, coordRobot)
				}
			} else {
				coordSanta.y += 1
				if !slices.Contains(visited, coordSanta) {
					visited = append(visited, coordSanta)
				}
			}
		case 'v':
			if i%2 == 0 {
				coordRobot.y -= 1
				if !slices.Contains(visited, coordRobot) {
					visited = append(visited, coordRobot)
				}
			} else {
				coordSanta.y -= 1
				if !slices.Contains(visited, coordSanta) {
					visited = append(visited, coordSanta)
				}
			}
		case '>':
			if i%2 == 0 {
				coordRobot.x += 1
				if !slices.Contains(visited, coordRobot) {
					visited = append(visited, coordRobot)
				}
			} else {
				coordSanta.x += 1
				if !slices.Contains(visited, coordSanta) {
					visited = append(visited, coordSanta)
				}
			}
		case '<':
			if i%2 == 0 {
				coordRobot.x -= 1
				if !slices.Contains(visited, coordRobot) {
					visited = append(visited, coordRobot)
				}
			} else {
				coordSanta.x -= 1
				if !slices.Contains(visited, coordSanta) {
					visited = append(visited, coordSanta)
				}
			}
		}
	}
	return len(visited)
}
