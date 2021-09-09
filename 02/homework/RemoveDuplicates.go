package homework

import "fmt"

// RemoveDuplicates function remove những phần tử bị trùng nhau trong mảng
func RemoveDuplicates(slice []int) []int {
	keys := make(map[int]bool)
	var result []int
	for _, entry := range slice {
		if _, value := keys[entry]; !value {
			keys[entry] = true
			fmt.Println(keys)
			result = append(slice, entry)
		}
	}
	return result
}
