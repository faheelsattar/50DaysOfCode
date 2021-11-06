package main

import (
	"fmt"
	"sort"
)

func heightChecker(heights []int, expected []int) int {
	sort.Ints(expected)
	count := 0
	for i := 0; i < len(heights); i++ {
		if heights[i] != expected[i] {
			count += 1
		}
	}
	fmt.Println("heigts expected", heights, expected)
	return count
}

func thirdMax(array []int) int {
	sort.Ints(array)
	if len(array) >= 3 {
		return array[len(array)-3]
	}
	return array[len(array)-1]
}

func disappearedArray(array []int) {
	sort.Ints(array)
	var hlder int = 0
	for i := 0; i < len(array); i++ {
		if i+1 < len(array) {
			if array[i] == array[i+1] || array[i]+1 == array[i+1] {

			} else {
				hlder = array[i] + 1
				fmt.Println(array[i] + 1)
			}

			if hlder+1 != array[i]+1 {

			}
		}
	}
}

func main() {
	heights := []int{1, 1, 4, 2, 1, 3}
	expected := []int{1, 1, 4, 2, 1, 3}
	fmt.Println("height checker", heightChecker(heights, expected))
	nums := []int{2, 2, 3, 1}
	fmt.Println("third max", thirdMax(nums))
	numArr := []int{4, 3, 2, 7, 8, 2, 3, 1}
	disappearedArray(numArr)
}
