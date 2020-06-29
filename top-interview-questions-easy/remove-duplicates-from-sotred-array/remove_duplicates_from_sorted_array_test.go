package remove_duplicates_from_sotred_array

import "testing"

func TestRemoveDuplicates(t *testing.T) {
	nums := []int{0, 1, 1, 2, 2, 3, 3, 3, 3, 3, 4}
	res := removeDuplicates(nums)
	t.Logf("result is %d, nums is %v", res, nums)

	if res == 5 {
		for i := 0; i < res; i++ {
			if nums[i] != i {
				t.Error("Wrong")
				return
			}
		}
	} else {
		t.Error("Wrong")
		return
	}
	t.Log("Correct")
}
