package _2

func preArr(arr []int) (sum []int) {
	sum = make([]int, len(arr))

	for i := range sum {
		if i == 0 {
			sum[i] = arr[i]
			continue
		}

		sum[i] = sum[i-1] + arr[i]
	}

	return
}

func getSum(sum []int, l, r int) int {
	if l == 0 {
		return sum[r]
	}

	return sum[r] - sum[l-1]
}
