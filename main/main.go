// https://www.hackerrank.com/challenges/insertionsort2/problem

package main

import "fmt"

func printArray(arr []int32) {
	for _, val := range arr {
		fmt.Printf("%d ", val)
	}

	fmt.Println()
}

func insertionSort2(n int32, arr []int32) {
	for i := int32(1); i < n; i++ {
		var temp = arr[i]

		for j := i - 1; j >= 0; j-- {
			if temp < arr[j] {
				arr[j+1] = arr[j]

				if j == 0 {
					arr[0] = temp
				}
			} else {
				if j < i-1 {
					arr[j+1] = temp
				}
				break
			}
		}

		printArray(arr)
	}
}

func main() {
	insertionSort2(7, []int32{3, 4, 7, 5, 6, 2, 1})
	insertionSort2(5, []int32{2, 4, 6, 8, 3})
	insertionSort2(10, []int32{2, 3, 4, 5, 6, 7, 8, 9, 10, 1})
}
