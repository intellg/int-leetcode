package single_number

import "testing"

type singleData struct {
	nums   []int
	result int
}

var testItems = []singleData{
	{
		nums:   []int{1, 2, 3, 1, 2, 3, 4},
		result: 4,
	},
	{
		nums:   []int{1, 2, 3, 1, 2, 3, 1},
		result: 1,
	},
	{
		nums:   []int{1, 2, 3, 1, 2, 3, 5, 6},
		result: 5,
	},
}

func TestSingleNumber(t *testing.T) {
	for _, testItem := range testItems {
		result := singleNumber(testItem.nums)
		if result != testItem.result {
			t.Errorf("Wrong for %v\n", testItem.nums)
			return
		}
	}
	t.Log("Correct")
}
