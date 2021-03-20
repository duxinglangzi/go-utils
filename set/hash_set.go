package set

import (
	"reflect"
	"sync"
)

type HashSet struct {
	set map[interface{}]bool
	sync.RWMutex
}

// 创建一个新对象
func New() *HashSet {
	return &HashSet{
		set: map[interface{}]bool{},
	}
}

// add element
func (s *HashSet) Add(ele interface{}) {
	s.Lock()
	defer s.Unlock()
	if len(s.set) != 0 {
		for e := range s.set {
			if reflect.TypeOf(e) == reflect.TypeOf(ele) {
				s.set[ele] = true
				break
			}else{
				panic("element is invalid, please check element type. ")
			}
		}
		
	}else{
		s.set[ele] = true
	}
}

// remove element
func (s *HashSet) Remove(e interface{}) {
	s.Lock()
	defer s.Unlock()
	delete(s.set, e)
}

// 是否包含
func (s *HashSet) Contains(e interface{}) bool {
	s.RLock()
	defer s.RUnlock()
	_, ok := s.set[e]
	return ok
}

// 长度
func (s *HashSet) Len() int {
	return len(s.set)
}

// 清空
func (s *HashSet) Clear() {
	s.Lock()
	defer s.Unlock()
	s.set = map[interface{}]bool{}
}

// 非空
func (s *HashSet) IsEmpty() bool {
	if s.Len() == 0 {
		return true
	}
	return false
}

// to list
func (s *HashSet) ToList() []interface{} {
	s.RLock()
	defer s.RUnlock()
	var list []interface{}
	for ele := range s.set {
		list = append(list, ele)
	}
	return list
}


