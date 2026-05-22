func calPoints(operations []string) int {
	out := []int{}

	for _, c := range operations {

		if c == "+" {
			n := out[len(out)-1] + out[len(out)-2]
			out = append(out, n)

		} else if c == "D" {
			n := 2 * out[len(out)-1]
			out = append(out, n)

		} else if c == "C" {
			out = out[:len(out)-1]

		} else {
			n, _ := strconv.Atoi(c)
			out = append(out, n)
		}
	}

	sum := 0
	for _, v := range out {
		sum += v
	}

	return sum
}