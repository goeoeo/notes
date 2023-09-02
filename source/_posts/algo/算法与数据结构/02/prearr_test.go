package _2

import (
	"fmt"
)

func ExamplePreArr() {
	arr := []int{1, 2, 3, 4, 5}
	sum := preArr(arr)

	r := getSum(sum, 2, 4)
	fmt.Println(r)
	//Output:12
}

func ExampleExits() {
	arr := []int{1, 2, 3, 4, 5}

	r := Exits(arr, 3)
	fmt.Println(r)
	//Output:true

}
