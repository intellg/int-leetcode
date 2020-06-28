package top_interview_questions_easy

import (
	"testing"
)

type testData struct {
	prices []int
	result int
}

var testDataList []testData

func init() {
	testDataList = []testData{
		{[]int{}, 0},
		{[]int{7, 1, 5, 3, 6, 4}, 7},
		{[]int{1, 2, 3, 4, 5}, 4},
		{[]int{7, 6, 5, 4, 3, 2, 1}, 0},
		{[]int{7, 1, 5, 3, 6, 4, 8}, 11},
		{[]int{7, 1, 5, 3, 6, 4, 8, 3, 7, 5, 3, 6, 8, 9, 8, 6, 5, 6, 8, 9, 7, 6, 5, 4, 3, 3, 3, 4, 6, 3}, 142},
	}
}

func TestMaxProfit(t *testing.T) {
	for _, testItem := range testDataList {
		result := maxProfit(testItem.prices)
		if result == testItem.result {
			t.Logf("correct result: %d", result)
		} else {
			t.Errorf("invalid result: expect %d but get %d\n", testItem.result, result)
		}
	}
}
