package datstr

// ShellSort is to sort an array
func ShellSort(Arr []int) []int {
	n := len(Arr)
	gap := n / 2
	for gap >= 1 {
		ssort(Arr, gap)
		gap = fngap(gap)
	}
	return Arr
}

func ssort(Arr []int, gap int) []int {

	for x := range Arr {
		for {
			if x+gap > len(Arr)-1 {
				break
			}

			if Arr[x] > Arr[x+gap] {
				Arr[x], Arr[x+gap] = Arr[x+gap], Arr[x]
			}
			x = x + gap

		}
	}
	return Arr
}

func fngap(n int) int {
	return n / 2
}
