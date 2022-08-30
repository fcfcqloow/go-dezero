package util

import "fmt"

type Set interface {
	fmt.Stringer
	Add(values ...interface{})
	Get(idx int) interface{}
	Find(key interface{}) interface{}
	Contains(key interface{}) bool
	Values() []interface{}
}

type set struct {
	keys   []interface{}
	values map[interface{}]interface{}
}

func NewSet(vs ...interface{}) Set {
	values := map[interface{}]interface{}{}
	keys := []interface{}{}
	for _, v := range vs {
		if _, ok := values[v]; !ok {
			keys = append(keys, v)
			values[v] = v
		}
	}
	return &set{
		values: values,
		keys:   keys,
	}
}

func (s *set) Add(values ...interface{}) {
	for _, v := range values {
		if _, ok := s.values[v]; !ok {
			s.keys = append(s.keys, v)
			s.values[v] = v
		}
	}
}

func (s *set) Contains(key interface{}) bool {
	_, ok := s.values[key]
	return ok
}

func (s *set) Get(idx int) interface{} {
	return s.keys[idx]
}

func (s *set) Find(key interface{}) interface{} {
	return s.values[key]
}

func (s *set) Values() []interface{} {
	return s.keys
}

func (s *set) String() string {
	return fmt.Sprintf("keys: %v, values: %v", s.keys, s.values)
}
