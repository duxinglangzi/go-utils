package set

import (
	"testing"
)

func TestHashSet_Add(t *testing.T) {
	set := New()
	set.Add("1")
	set.Add("3")
	t.Log(set.Len())
}

func TestHashSet_Add2(t *testing.T) {
	set := New()
	set.Add(1)
	set.Add(2)
	set.Add(4)
	t.Log(set.Len())
	
	set.Add("4") //  panic: element is invalid, please check element type.
	
}

func TestHashSet_Clear(t *testing.T) {
	set := New()
	set.Add(1)
	set.Add(2)
	set.Add(4)
	t.Log(set.Len()) // 3
	set.Clear()
	t.Log(set.Len()) // 0
}

func TestHashSet_Remove(t *testing.T) {
	set := New()
	set.Add(3456)
	set.Add(19573)
	t.Log(set.ToList()) // [3456 19573]
	
	set.Remove(19573)
	
	t.Log(set.ToList()) //  [3456]
	
	
}

