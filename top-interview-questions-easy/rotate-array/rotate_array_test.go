package rotate_array

import "testing"

type rotateArrayData struct {
	nums   []int
	k      int
	result []int
}

var rotateArray = []rotateArrayData{
	{
		nums:   []int{1},
		k:      0,
		result: []int{1},
	},
	{
		nums:   []int{1, 2, 3, 4, 5, 6, 7, 8, 9},
		k:      3,
		result: []int{7, 8, 9, 1, 2, 3, 4, 5, 6},
	},
	{
		nums:   []int{1, 2},
		k:      2,
		result: []int{1, 2},
	},
	{
		nums:   []int{1, 2, 3, 4, 5, 6, 7},
		k:      12,
		result: []int{3, 4, 5, 6, 7, 1, 2},
	},
}

func TestRotate(t *testing.T) {
	for _, testItem := range rotateArray {
		rotate(testItem.nums, testItem.k)
		for i := 0; i < len(testItem.nums); i++ {
			if testItem.nums[i] != testItem.result[i] {
				t.Errorf("Wrong for %v\n", testItem.nums)
			}
		}
	}
	t.Log("Correct")
}
