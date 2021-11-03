package main

import (
	"fmt"
	"sort"
)

//q1
func duplicateZeros(numbers []int) []int {
	var answer []int
	var i int
	for i = 0; i < len(numbers); i++ {
		answer = append(answer, numbers[i])
		if numbers[i] == 0 {
			answer = append(answer, 0)
			numbers = numbers[:len(numbers)-1]
		}
	}
	fmt.Println("num", numbers)
	return answer
}

func merge(numArr1 []int, k int, numArr2 []int, j int) []int {
	var i int
	for i = 0; i < len(numArr2); i++ {
		if numArr1[k+i] == 0 {
			numArr1[k+i] = numArr2[i]
		}
	}
	sort.Ints(numArr1)

	return numArr1
}

func removeElements(numArr []interface{}, value int) int {
	counter := 0
	for i := range numArr {
		if numArr[i] == value {
			numArr[i] = "_"
			counter += 1
		}
	}
	return counter
}

//implemented via recurssion
func removeDuplicates(numArr []interface{}, currentIndex int, counter int) int {
	//base case
	var nextvalue interface{}
	if currentIndex == len(numArr) {
		return counter
	}
	if currentIndex+1 < len(numArr) {
		nextvalue = numArr[currentIndex+1]
	}
	if numArr[currentIndex] == nextvalue {
		numArr[currentIndex] = "_"
	} else {
		counter += 1
	}
	return removeDuplicates(numArr, currentIndex+1, counter)
}

func main() {
	//q1
	var numbers = []int{1, 0, 2, 3, 0, 4, 5, 0}
	fmt.Println("duplicate zeros", duplicateZeros(numbers))

	//q2
	var num1Arr []int = []int{1, 2, 3, 0, 0, 0}
	var numArr2 []int = []int{2, 5, 6}
	fmt.Println("merged sorted arr", merge(num1Arr, 3, numArr2, 3))

	//q3
	numArr := []interface{}{3, 2, 2, 3}
	fmt.Println("Counter", removeElements(numArr, 3))

	//q4
	numArr1 := []interface{}{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}
	fmt.Println("total unique values", removeDuplicates(numArr1, 0, 0))
}
