package main

import (
	"log"
)

func main() {
	in := []int{1, 2, 3, 4, 5, 6, 7}
	rotate(in, 1)
	for i := range in {
		log.Println(in[i])
	}
}

func rotate(nums []int, k int) {
	l := len(nums)
	if l <= 0 {
		return
	}
	for k != 0 {
		var temp, i int
		temp = nums[0]
		for i < l-1 {
			i++
			nums[0] = nums[i]
			nums[i] = temp
			temp = nums[0]
		}
		nums[0] = temp
		k--
	}
}

/* func rotate(nums []int, k int) {
	l := len(nums)
	a := make([]int, l)
	for i := range nums {
		n := (i + k) % l
		a[n] = nums[i]
	}

	for i := range nums {
		nums[i] = a[i]
	}
} */
