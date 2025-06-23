package main

func restoreString(s string, indices []int) string {
	result := make([]byte, len(s))

	for i, v := range indices {
		result[v] = s[i] //result[4] = s[0]
	}
	return string(result)
}

//func main() {
//	result := restoreString("codeleet", []int{4, 5, 6, 7, 0, 2, 1, 3})
//
//	fmt.Println(result)
//}

// O(n)
// O(n)
