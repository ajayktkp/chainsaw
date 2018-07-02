package main

/* func main() {
	in := []int{1, 2, 3, 4, 5, 6, 7}
	rotate(in, 2)
	log.Println("----------------------------------")
	for i := range in {
		log.Println(in[i])
	}
} */

func rotate(nums []int, k int) {
	l := len(nums)
	if l <= 0 {
		return
	}
	k = k % l
	var i int
	for i < l-1 {
		n := idx(i, k, l)
		temp := nums[i]
		nums[i] = nums[n]
		nums[n] = temp
		i++
	}

}

func idx(curr, k, l int) int {
	n := (curr - k + l) % l
	if curr < k {
		return n
	}

	for n < k || n < curr {
		n = idx(n, k, l)
	}
	return n
}

/* func rotate(nums []int, k int) {
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
} */

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
