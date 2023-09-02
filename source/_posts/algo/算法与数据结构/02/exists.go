package _2

func Exits(arr []int, num int) bool {
	if len(arr) == 0 {
		return false
	}

	l := 0
	r := len(arr) - 1
	m := 0
	for l <= r {
		m = l + (r-l)>>1
		if arr[m] == num {
			return true
		} else if arr[m] > num {
			r = m - 1
		} else {
			l = m + 1
		}
	}

	return false
}
