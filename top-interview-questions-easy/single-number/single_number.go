// https://leetcode.com/explore/featured/card/top-interview-questions-easy/92/array/549/
package single_number

func singleNumber(nums []int) int {
	keys := make(map[int]int)
	for _, num := range nums {
		if _, ok := keys[num]; ok {
			delete(keys, num)
		} else {
			keys[num] = 0
		}
	}
	for key, _ := range keys {
		return key
	}
	return -1
}
