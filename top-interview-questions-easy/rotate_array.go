// https://leetcode.com/explore/featured/card/top-interview-questions-easy/92/array/646/
package top_interview_questions_easy

func rotate(nums []int, k int) {
	if k == 0 {
		return
	}

	a := len(nums)
	b := k
	for a != b {
		if a > b {
			a = a - b
		} else {
			b = b - a
		}
	}

	var extra int
	for j := 0; j < a; j++ {
		i := j
		p := -1
		for {
			if p == -1 {
				extra = nums[i]
			} else {
				nums[p] = nums[i]
			}
			p = i

			i -= k
			for i < 0 {
				i += len(nums)
			}
			if i == j {
				break
			}
		}
		nums[p] = extra
	}
}

const length = 1000000

func test1() int {
	array := make([]int, length)
	sum := 0
	for i := 0; i < length; i++ {
		sum += len(array)
	}
	return sum
}

func test2() int {
	array := make([]int, length)
	length := len(array)
	sum := 0
	for i := 0; i < length; i++ {
		sum += length
	}
	return sum
}

func test3() int {
	array := make([]int, length)
	sum := 0
	for i := 0; i < len(array); i++ {
		sum += array[i]
	}
	return sum
}

func test4() int {
	array := make([]int, length)
	sum := 0
	for i := 0; i < length; i++ {
		sum += array[i]
	}
	return sum
}

func test5() int {
	array := make([]int, length)
	sum := 0
	for _, item := range array {
		sum += item
	}
	return sum
}
