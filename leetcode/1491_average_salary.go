package main

import "fmt"

func average(salary []int) float64 {
	min := salary[0]
	max := salary[0]
	total := 0
	n := len(salary)

	for i := 0; i < len(salary); i++ {
		if min > salary[i] {
			min = salary[i]
		}

		if max < salary[i] {
			max = salary[i]
		}

		total += salary[i]
	}

	total = total - max - min
	result := float64(total) / float64(n-2)
	return result
}

func main() {
	salary := average([]int{1000, 2000, 3000, 4000})
	fmt.Println(salary)
}
