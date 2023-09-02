package sort

func InsertSort(arr []int) {
	if arr == nil || len(arr) < 2 {
		return
	}

	n := len(arr)
	for i := 0; i < n; i++ {
		for j := i - 1; j >= 0 && arr[j] < arr[j+1]; j-- {
			arr[j], arr[j+1] = arr[j+1], arr[j]
		}
	}
}
