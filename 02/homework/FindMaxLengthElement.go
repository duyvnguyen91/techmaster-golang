package homework

// FindMaxLengthElement Lọc ra các phần tử có độ dài lớn nhất
func FindMaxLengthElement(slice []string) []string {
	max := 1
	var result []string
	for _, v := range slice {
		if len(v) < max {
			continue
		}
		if len(v) > max {
			max = len(v)
			result = result[:0]
		}
		result = append(result, v)
	}
	return result
}