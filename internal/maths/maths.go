package maths

func SumSlice(slice []int) int {
	r := 0
	for _, n := range slice {
		r += n
	}
	return r
}

func AbsInt(i int) int {
	if i < 0 {
		return i * -1
	}
	return i
}
