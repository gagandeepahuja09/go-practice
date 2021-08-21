package sum

import (
	"testing"
)

func TestInt(t *testing.T) {
	tt := []struct {
		name    string
		numbers []int
		sum     int
	}{
		{"1 to 5", []int{1, 2, 3, 4, 5}, 15},
		{"nil input", nil, 1},
		{"1 and -1", []int{1, -1}, 0},
	}

	for _, tc := range tt {
		result := Ints(tc.numbers...)
		// Don't know why but assert is not working here
		// assert.Equal(t, tc.sum, result)
		if tc.sum != result {
			t.Errorf("Test %s failed expected %d found %d", tc.name, tc.sum, result)
		}
	}
}
