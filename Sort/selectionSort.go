package datstr

//SelectionSort is for sorting an array
func SelectionSort(Arr []int) []int {

	for i := 0; i < len(Arr)-2; i++ {
		iMin := i
		for j := i + 1; j <= len(Arr)-1; j++ {
			if Arr[iMin] > Arr[j] {
				iMin = j
			}
		}
		Arr[i], Arr[iMin] = Arr[iMin], Arr[i]
	}
	return Arr
}
