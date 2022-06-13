package main

import "fmt"

func main() {
	var count int
	var arr = []int{12, 4, 56, 34, 34, 4, 12}

	for i := 0; i < len(arr); i++ {
		count = 0
		//	fmt.Println(arr[i])

		for j := 0; j < len(arr); j++ {
			//fmt.Println(arr[j])
			if i != j {
				if arr[i] == arr[j] {

					count = count + 1
				}
			}
		}

		if count == 0 {
			fmt.Println(arr[i])
		}

	}

}
