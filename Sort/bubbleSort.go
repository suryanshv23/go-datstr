package datstr

//BubbleSort is for bubble sorting an array
func BubbleSort(Arr []int) []int {
	for i := 0; i <= len(Arr)-2; i++ {
		for j := i + 1; j <= len(Arr)-1; j++ {
			if Arr[i] > Arr[j] {
				Arr[i], Arr[j] = Arr[j], Arr[i]
			}
		}
	}
	return Arr
}
