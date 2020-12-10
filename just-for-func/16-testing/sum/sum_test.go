package sum

import (
	"testing"
)

func TestInts(t *testing.T) {
	testingTable := []struct {
		name    string
		numbers []int
		sum     int
	}{
		{"1-to-5", []int{1, 2, 3, 4, 5}, 15},
		{"no numbers", []int{}, 0},
		{"1 and -1", []int{1, -1}, 0},
		{"nil", nil, 0},
	}

	for _, testCase := range testingTable {
		t.Run(testCase.name, func(t *testing.T) {
			result := Ints(testCase.numbers...)
			if result != testCase.sum {
				t.Fatalf("Sum of %v must be %v but we got %v\n", testCase.name, testCase.sum, result)
			}
		})
	}
}
