package homework

import "sort"

// Max2Numbers Tìm ra số lớn thứ 2 trong mảng
func Max2Numbers(numbers []int) int {
	sort.Slice(numbers, func(i, j int) bool {
		return numbers[i] > numbers[j]
	} )
	return numbers[1]
}
