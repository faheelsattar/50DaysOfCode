// [1,0,1,1,0,0,0,0,1,1,1,1,0,1,1,1,0]
package main

import (
	"fmt"
)

func findMax(a []int) int {
	var max int = a[0]
	for _, value := range a {
		if value > max {
			max = value
		}
	}
	return max
}

//question 1 ==> implemented via recursion
func consecutiveOnes(arr []int, index int, max int, repeat bool, answer []int) []int {
	// base case
	if index == len(arr) {
		return answer
	}
	if arr[index] == 1 {
		if repeat {
			max = max + 1
			answer = append(answer, max)
		} else {
			answer = append(answer, max)
			max = 1
		}
		if index+1 < len(arr) && arr[index+1] == 1 {
			repeat = true
		} else {
			repeat = false
		}
		return consecutiveOnes(arr, index+1, max, repeat, answer)
	} else {
		return consecutiveOnes(arr, index+1, max, false, answer)
	}
}

//question 2
func evenNumCounter(numArr []int) int {
	var counter int = 0
	for _, data := range numArr {
		if data%2 == 0 {
			counter += 1
		}
	}
	return counter
}

//question 3
func squaredArray(numArr []int) []int {
	for i := range numArr {
		numArr[i] = numArr[i] * numArr[i]
	}
	return numArr
}

func main() {

	//q1
	var numArray []int = []int{1, 1, 0, 1, 0, 1, 0, 1, 0, 0, 1, 1, 1, 1, 1, 1, 0, 1, 0, 0, 0, 0, 1, 1, 1, 0}
	var answer []int
	var countOnes = consecutiveOnes(numArray, 0, 0, false, answer)
	var max = findMax(countOnes)
	fmt.Println("max consecutive 1's: ", max)

	//q2
	var numbers []int = []int{21, 12, 12, 12, 7, 3, 4}
	var evenCounter = evenNumCounter(numbers)
	fmt.Println("total even numbers", evenCounter)

	//q3
	var numbers1 []int = []int{1, 1, 1, 2, 3, 4, 4, 4, 5, 6}
	fmt.Println("squared sequence", squaredArray(numbers1))
}
