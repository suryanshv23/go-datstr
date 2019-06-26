package datstr

import "fmt"

//QuickSort is for sorting an array
func QuickSort(Arr []int) []int {
	if len(Arr) < 2 {
		return Arr
	}
	start, index := 0, 0
	pivot := len(Arr) - 1

	for i := start; i <= len(Arr)-1; i++ {
		if Arr[i] < Arr[pivot] {
			fmt.Println(Arr, Arr[index], Arr[i])

			Arr[index], Arr[i] = Arr[i], Arr[index]
			index++
		}
	}
	Arr[pivot], Arr[index] = Arr[index], Arr[pivot]
	fmt.Println(Arr[pivot], Arr[index])

	QuickSort(Arr[:index])
	QuickSort(Arr[index+1:])
	return Arr
}
