package mergesort

import (
	"math/rand"
	"reflect"
	"testing"
	"time"
)

func TestMergeSort(t *testing.T) {
	testCases := []struct {
		test   string
		arg    []int
		wanted []int
	}{
		{
			"even amount of numbers",
			[]int{3, 8, 6, 9, 1, 5, 2, 0, 7, 4},
			[]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9},
		},
		{
			"odd amount of numbers",
			[]int{3, 8, 6, 1, 5, 2, 0, 7, 4},
			[]int{0, 1, 2, 3, 4, 5, 6, 7, 8},
		},
		{
			"no numbers",
			[]int{},
			[]int{},
		},
		{
			"one number",
			[]int{1},
			[]int{1},
		},
	}

	for _, testCase := range testCases {
		got := Sort(testCase.arg)

		if !reflect.DeepEqual(testCase.wanted, got) {
			t.Fatalf("%s - Wanted %+v, but got %+v", testCase.test, testCase.wanted, got)
		}
	}
}

func BenchmarkMergeSort(b *testing.B) {
	for i := 0; i < b.N; i++ {
		rand.Seed(time.Now().Unix())
		arg := rand.Perm(1000)
		Sort(arg)
	}
}
