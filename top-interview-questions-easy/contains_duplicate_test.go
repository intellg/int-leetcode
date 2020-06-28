package top_interview_questions_easy

import "testing"

type duplicateData struct {
	nums   []int
	result bool
}

func TestContainsDuplicate(t *testing.T) {
	testItems := []duplicateData{
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

	for _, testItem := range testItems {
		result := containsDuplicate(testItem.nums)
		if result != testItem.result {
			t.Errorf("Wrong for %v", testItem.nums)
			return
		}
	}
	t.Log("Correct")
}
