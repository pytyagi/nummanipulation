package num

// ArraySumInt return sum of array elements
func ArraySumInt(a ...int64) int64 {
	var sum int64
	for i := 0; i < len(a); i++ {
		sum += a[i]
	}
	return sum
}
