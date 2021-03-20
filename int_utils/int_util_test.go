package int_utils

import "testing"

func TestContains(t *testing.T) {
	slice := []int{1,2,3,4}
	t.Log(Contains(slice, 3))
	t.Log(Contains(slice, 8))
}
