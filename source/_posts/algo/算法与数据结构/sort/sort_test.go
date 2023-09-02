package sort

import "fmt"

func ExampleSelectSort() {
	arr := []int{4, 3, 2, 1, 5}

	SelectSort(arr)
	fmt.Println(arr)
	//Output:[1 2 3 4 5]
}

func ExampleBubblingSort() {
	arr := []int{4, 3, 2, 1, 5}

	BubblingSort(arr)
	fmt.Println(arr)
	//Output:[1 2 3 4 5]
}

func ExampleInsertSort() {
	arr := []int{4, 3, 2, 1, 5}

	InsertSort(arr)
	fmt.Println(arr)
	//Output:[1 2 3 4 5]
}
