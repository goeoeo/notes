package sort

// SelectSort 选择排序
func SelectSort(arr []int) {
	if arr == nil || len(arr) < 2 {
		return
	}

	n := len(arr)
	for i := 0; i < n-1; i++ {
		minIndex := i

		//找出最小的位置
		for j := i + 1; j < n; j++ {
			if arr[j] < arr[minIndex] {
				minIndex = j
			}
		}

		//与当前位置交换
		arr[i], arr[minIndex] = arr[minIndex], arr[i]
	}
}
