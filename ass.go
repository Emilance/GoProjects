package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func sortPartition(_slice []int, ch chan []int) {
	sort.Ints(_slice)
	ch <- _slice
}

func merge(first, second []int) []int {
	resultant := []int{}
	f := len(first)
	s := len(second)
	i := 0
	j := 0

	for i < f && j < s {
		if first[i] < second[j] {
			resultant = append(resultant, first[i])
			i++
		} else {
			resultant = append(resultant, second[j])
			j++
		}
	}

	for ; i < f; i++ {
		resultant = append(resultant, first[i])
	}
	for ; j < s; j++ {
		resultant = append(resultant, second[j])
	}

	return resultant
}

func main() {
	fmt.Println("Enter space separated array elements ")
	reader := bufio.NewReader(os.Stdin)
	input, _, err := reader.ReadLine()
	_input := strings.Trim(string(input), " ")
	ch := make(chan []int, 4)
	var array []int

	if err == nil {
		stringArray := strings.Split(_input, " ")
		len := len(stringArray)
		for k := 0; k < len; k++ {
			val, err2 := strconv.Atoi(stringArray[k])
			if err2 != nil {
				fmt.Println("Enter valid integer values")
				return
			}
			array = append(array, val)
		}
		if len < 4 {
			fmt.Println("Enter more than 4 elements in the array")
		} else {
			factor := int(len / 4)
			rem := len % 4

			start := 0
			end := factor
			if rem != 0 {
				end = end + 1
				rem = rem - 1
			}
			go sortPartition(array[start:end], ch)

			start = end
			end = end + factor
			if rem != 0 {
				end = end + 1
				rem = rem - 1
			}
			go sortPartition(array[start:end], ch)

			start = end
			end = end + factor
			if rem != 0 {
				end = end + 1
				rem = rem - 1
			}
			go sortPartition(array[start:end], ch)

			start = end
			end = end + factor
			if rem != 0 {
				end = end + 1
				rem = rem - 1
			}
			go sortPartition(array[start:end], ch)

			part1 := <-ch
			part2 := <-ch
			part3 := <-ch
			part4 := <-ch

			fmt.Println("\nSorted Part1 of Array : ", part1)
			fmt.Println("Sorted Part2 of Array : ", part2)
			fmt.Println("Sorted Part3 of Array : ", part3)
			fmt.Println("Sorted Part4 of Array : ", part4)

			result := merge(part1, part2)
			result = merge(result, part3)
			result = merge(result, part4)
			fmt.Println("Sorted Array : ", result)
		}
	} else {
		fmt.Println("Error occurred. ", err)
	}
}
