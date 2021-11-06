package main

import "fmt"

// [10,2,5,3]
//implemented via recurrsion
func doubleExists(numbers []int, index int) bool {
	if index == len(numbers) {
		return false
	}
	for i := 0; i < len(numbers); i++ {
		if numbers[i] == numbers[index]*2 {
			return true
		}
	}
	return doubleExists(numbers, index+1)
}

func mountainArray(arr []int) bool {
	mid := len(arr) % 2
	if mid == 0 {
		mid = len(arr) / 2
	} else {
		mid = (len(arr) / 2) + 1
	}
	if len(arr) < 3 {
		return false
	}

	fmt.Println("mid", 0, mid, len(arr))
	for i := 0; i < mid; i++ {
		if i+1 < mid {
			if arr[i] >= arr[i+1] {
				return false
			}
		}
	}
	for i := mid; i < len(arr); i++ {
		if i+1 < len(arr) {
			if arr[i] <= arr[i+1] {
				return false
			}
		}
	}
	return true
}

// func replaceElements(arr []int, index int) []int {
// 	if index == len(arr) {
// 		fmt.Println("answer", arr)
// 		return arr
// 	}
// 	for i := index; i < len(arr); i++ {
// 		if i+1 < len(arr) {
// 			if arr[index] < arr[i+1] {
// 				arr[index] = arr[i+1]
// 				fmt.Println(arr[i+1])
// 			} else {
// 				arr[index] = arr[i]
// 				fmt.Println(arr[index])
// 			}
// 		}
// 	}
// 	return replaceElements(arr, index+1)
// }

//[0,1,0,3,12]

func moveZeroes(arr []int) []int {
	count := 0
	for i := 0; i < len(arr); i++ {
		if arr[i] != 0 {
			arr[count] = arr[i]
			count += 1
		}
	}
	for count < len(arr) {
		arr[count] = 0
		count += 1
	}
	return arr
}

func sortByParity(arr []int) []int {
	var answer []int
	for i := 0; i < len(arr); i++ {
		if arr[i]%2 == 0 {
			answer = append(answer, arr[i])
		}
	}
	for i := 0; i < len(arr); i++ {
		if arr[i]%2 == 1 {
			answer = append(answer, arr[i])
		}
	}

	return answer
}

func removeElement(arr []interface{}, val int) []interface{} {
	count := 0
	for i := 0; i < len(arr); i++ {
		if arr[i] != val {
			arr[count] = arr[i]
			count += 1
		}
	}
	for count < len(arr) {
		arr[count] = "_"
		count += 1
	}
	return arr

}
func main() {
	// numbers := []int{10, 9, 5, 1}
	// println("double exists", doubleExists(numbers, 0))
	// arr := []int{2, 1}
	// fmt.Println("mountain array", mountainArray(arr))

	array := []int{0, 1, 0, 3, 12}
	fmt.Println("moved zeroes array", moveZeroes(array))

	arr := []int{3, 1, 2, 4}
	fmt.Println("sort by parity", sortByParity(arr))
	numArr1 := []interface{}{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}
	fmt.Println("remove elements", removeElement(numArr1, 3))
}
