package main

import "fmt"

func sorting(arr []int) []int {
	for i := 0; i < len(arr); i++ {
		for j := i + 1; j < len(arr); j++ {
			if arr[i] > arr[j] {
				arr[i], arr[j] = arr[j], arr[i]
			}
		}
	}
	return arr

}

func main() {
	arr := []int{5, 2, 8, 1, 15}
	fmt.Println(sorting(arr))

}
