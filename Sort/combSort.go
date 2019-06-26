package datstr

//CombSort is for faster sorting of an array
func CombSort(Arr []int) []int {
	n := len(Arr)
	gap := n - 1

	for {
		if gap < 1 {
			gap = 1
		}
		csort(Arr, gap)
		if gap == 1 {
			break
		}
		gap = gapfn(float64(gap))
	}
	return Arr
}

func csort(Arr []int, gap int) []int {
	for i := 0; i+gap <= len(Arr)-1; i++ {

		if Arr[i] > Arr[i+gap] {
			Arr[i], Arr[i+gap] = Arr[i+gap], Arr[i]
		}
	}
	return Arr
}

func gapfn(n float64) int {
	return int(n / 1.3)
}
