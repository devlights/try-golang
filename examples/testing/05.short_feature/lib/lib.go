package lib

func Add(x, y int) int {
	return x + y
}

func Sum(num ...int8) int {
	var (
		total int
	)

	for _, x := range num {
		total += int(x)
	}

	return total
}
