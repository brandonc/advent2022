package maths

func SumSlice(slice []int) int {
	r := 0
	for _, n := range slice {
		r += n
	}
	return r
}
