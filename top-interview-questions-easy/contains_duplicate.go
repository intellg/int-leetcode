// https://leetcode.com/explore/featured/card/top-interview-questions-easy/92/array/578/
package top_interview_questions_easy

func containsDuplicate(nums []int) bool {
	keys := make(map[int]int)
	for _, num := range nums {
		if _, ok := keys[num]; ok {
			return true
		} else {
			keys[num] = 0
		}
	}
	return false
}
