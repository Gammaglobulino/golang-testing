package andrea

func MySum(v ...int) int {
	sum := 0
	for _, x := range v {
		sum += x
	}
	return sum
}
