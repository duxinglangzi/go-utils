package string_utils

import "testing"

func TestContains(t *testing.T) {
	t.Log("testing   hahahahhahahah")
	slice := []string{"1","2","3"}
	contains := Contains(slice, "2")
	t.Log(contains)
}
