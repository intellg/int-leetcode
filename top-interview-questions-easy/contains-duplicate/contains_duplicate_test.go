package contains_duplicate

import "testing"

type containsDuplicateData struct {
	nums   []int
	result bool
}

var containsDuplicateItems = []containsDuplicateData{
	{
		nums:   []int{1, 2, 3, 4, 5, 6, 7, 8},
		result: false,
	},
	{
		nums:   []int{1, 2, 3, 4, 5, 6, 7, 1},
		result: true,
	},
	{
		nums:   []int{1, 1, 2, 3, 4, 5, 6, 7, 8},
		result: true,
	},
}

func TestContainsDuplicate(t *testing.T) {
	for _, testItem := range containsDuplicateItems {
		result := containsDuplicate(testItem.nums)
		if result != testItem.result {
			t.Errorf("Wrong for %v\n", testItem.nums)
			return
		}
	}
	t.Log("Correct")
}
