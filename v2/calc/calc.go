package calc

// Add sum two int parameters and return the result
func Add(args ...int) int {
	s := 0
	for _, v := range args {
		s += v
	}
	return s
}

// Sub Subtract two int parameters and return the result
func Sub(a, b int) int {
	return a - b
}
