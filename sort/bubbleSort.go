package sort

func BubbleSort(arr []int32) []int32 {
	for j := 0; j < len(arr); j++ {
		for i := 0; i < len(arr)-1; i++ {
			if arr[i] > arr[i+1] {
				arr[i], arr[i+1] = arr[i+1], arr[i]
			} else {
			}
		}
	}
	return arr
}
