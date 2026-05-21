func replaceElements(arr []int) []int {
	max:= -1
	for i:= len(arr) - 1;i >= 0 ; i -- {
		current := arr[i]
		arr[i] = max
		if current > max {
			max = current
		}
	}
	return arr
}
