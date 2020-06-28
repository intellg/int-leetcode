// https://leetcode.com/explore/featured/card/top-interview-questions-easy/92/array/727/
package top_interview_questions_easy

func removeDuplicates(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	target := 1
	previous := nums[0]
	for i := 1; i < len(nums); i++ {
		if nums[i] != previous {
			nums[target] = nums[i]
			target++
		}
		previous = nums[i]
	}
	return target
}
