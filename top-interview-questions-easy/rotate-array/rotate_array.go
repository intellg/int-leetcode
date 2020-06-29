// https://leetcode.com/explore/featured/card/top-interview-questions-easy/92/array/646/
package rotate_array

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
