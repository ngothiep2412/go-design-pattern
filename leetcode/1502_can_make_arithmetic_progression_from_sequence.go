package main

import "fmt"

func canMakeArithmeticProgression(arr []int) bool {
	min := arr[0]
	max := arr[0]
	N := len(arr)

	existNumber := make(map[int]bool)
	for i := 0; i < len(arr); i++ {
		if arr[i] < min {
			min = arr[i]
		}

		if arr[i] > max {
			max = arr[i]
		}

		existNumber[arr[i]] = true
	}

	totalDiff := max - min
	if totalDiff%(N-1) != 0 {
		return false
	}

	diff := totalDiff / (N - 1)

	for min < max {
		min += diff

		if _, isExist := existNumber[min]; !isExist {
			return false
		}

	}

	return true
}

func main() {
	result := canMakeArithmeticProgression([]int{1, 4, 3})

	fmt.Println(result)
}
