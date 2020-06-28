package top_interview_questions_easy

import "testing"

type rotateArrayData struct {
	nums   []int
	k      int
	result []int
}

func TestRotate(t *testing.T) {
	testItems := []rotateArrayData{
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

	for _, testItem := range testItems {
		rotate(testItem.nums, testItem.k)
		for i := 0; i < len(testItem.nums); i++ {
			if testItem.nums[i] != testItem.result[i] {
				t.Error("Wrong")
			}
		}
	}
	t.Log("Correct")
}

func BenchmarkTest1(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		test1()
	}
}

func BenchmarkTest2(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		test2()
	}
}

func BenchmarkTest3(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		test3()
	}
}

func BenchmarkTest4(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		test4()
	}
}

func BenchmarkTest5(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		test5()
	}
}
