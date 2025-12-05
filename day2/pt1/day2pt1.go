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

	id_ranges := strings.Split(string(dat), ",")
	invalid_id_sum := 0

	for _, id_range := range id_ranges {

		id_range := strings.Split(string(id_range), "-")

		start, err := strconv.Atoi(id_range[0])
		check(err)
		end, err := strconv.Atoi(id_range[1])
		check(err)

		for i := start; i <= end; i++ {
			i_str := strconv.Itoa(i)
			len_i := len(i_str)
			if i_str[:len_i/2] == i_str[len_i/2:] {
				invalid_id_sum += i
			}
		}

	}

	fmt.Println(invalid_id_sum)

}
