package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	dat, err := os.ReadFile("../input.txt")
	check(err)

	instructions_rows := strings.Split(string(dat), "\n")

	zeros_reached := 0
	position := 50

	for _, row := range instructions_rows {

		direction := row[0:1]
		steps, err := strconv.Atoi(row[1:])
		check(err)

		if direction == "L" {
			position -= steps % 100
			if position < 0 {
				position += 100
			}
		} else if direction == "R" {
			position += steps % 100
			if position > 99 {
				position -= 100
			}
		}
		
		if position == 0 {
			zeros_reached++
		}

	}

	fmt.Println(zeros_reached)

}
