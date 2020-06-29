// https://leetcode.com/explore/featured/card/top-interview-questions-easy/92/array/564/
package best_time_to_buy_and_sell_stock_ii

import "fmt"

func maxProfit(prices []int) (result int) {
	previous := 0
	train := make([][]int, 0)
	length := len(prices)
	for i := 1; i < length; i++ {
		if prices[i] <= prices[i-1] {
			if previous < i-1 {
				train = append(train, []int{previous, i - 1})
			}
			previous = i
		}
	}
	if previous < length-1 {
		train = append(train, []int{previous, length - 1})
	}

	for i := 0; i < len(train); i++ {
		result += prices[train[i][1]] - prices[train[i][0]]
	}

	fmt.Print("[")
	for j := 0; j < len(train); j++ {
		fmt.Printf("[%d %d]", prices[train[j][0]], prices[train[j][1]])
	}
	fmt.Println("]")
	return
}

func maxProfit2(prices []int) (result int) {
	profits := make([][]int, 0)
	for i := 0; i < len(prices); i++ {
		for j := i + 1; j < len(prices); j++ {
			if prices[i] < prices[j] {
				profits = append(profits, []int{i, j, prices[j] - prices[i]})
			}
		}
	}

	relations := make([][]int, 0)
	for i := 0; i < len(profits); i++ {
		for j := i + 1; j < len(profits); j++ {
			if profits[i][1] < profits[j][0] {
				for k := j; k < len(profits); k++ {
					relations = append(relations, []int{i, k})
				}
				break
			}
		}
	}

	trains := make([][]int, 0)
	for i := 0; i < len(profits); i++ {
		train := make([]int, 1)
		train[0] = i
		trains = append(trains, train)
	}
	for i := 0; i < len(trains); i++ {
		train := trains[i]
		for j := 0; j < len(relations); j++ {
			relation := relations[j]
			if train[len(train)-1] == relation[0] {
				newTrain := append(train, relations[j][1])
				trains = append(trains, newTrain)
			}
		}
	}

	var finalTrain []int
	for i := 0; i < len(trains); i++ {
		train := trains[i]
		sum := 0
		for j := 0; j < len(train); j++ {
			sum += profits[train[j]][2]
		}
		if sum > result {
			result = sum
			finalTrain = train
		}
	}

	fmt.Print("[")
	for j := 0; j < len(finalTrain); j++ {
		fmt.Printf("[%d %d]", prices[profits[finalTrain[j]][0]], prices[profits[finalTrain[j]][1]])
	}
	fmt.Println("]")
	return
}
