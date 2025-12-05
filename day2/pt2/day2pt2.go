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

			for seq_len := 1; seq_len <= len_i/2; seq_len++ {
				initial_seq := i_str[:seq_len]
				is_invalid := true

				if len_i%seq_len != 0 {
					continue
				}

				for chunk_num := range (len_i / seq_len) - 1 {
					if i_str[(1+chunk_num)*seq_len:(2+chunk_num)*seq_len] != initial_seq {
						is_invalid = false
						break
					}
				}

				if is_invalid {
					fmt.Println(i)
					invalid_id_sum += i
					break
				}
			}
		}

	}

	fmt.Println(invalid_id_sum)

}
