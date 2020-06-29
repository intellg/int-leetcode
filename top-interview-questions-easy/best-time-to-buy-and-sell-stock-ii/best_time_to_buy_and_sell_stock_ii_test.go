package best_time_to_buy_and_sell_stock_ii

import (
	"testing"
)

type priceData struct {
	prices []int
	result int
}

var priceItems = []priceData{
	{[]int{}, 0},
	{[]int{7, 1, 5, 3, 6, 4}, 7},
	{[]int{1, 2, 3, 4, 5}, 4},
	{[]int{7, 6, 5, 4, 3, 2, 1}, 0},
	{[]int{7, 1, 5, 3, 6, 4, 8}, 11},
	{[]int{7, 1, 5, 3, 6, 4, 8, 3, 7, 5, 3, 6, 8, 9, 8, 6, 5, 6, 8, 9, 7, 6, 5, 4, 3, 3, 3, 4, 6, 3}, 142},
}

func TestMaxProfit(t *testing.T) {
	for _, testItem := range priceItems {
		result := maxProfit(testItem.prices)
		if result != testItem.result {
			t.Errorf("invalid result: expect %d but get %d\n", testItem.result, result)
		}
	}
	t.Log("Correct")
}
