package ds

type Stack[T any] struct {
	items []T
}

func (s *Stack[T]) Push(item T) {
	s.items = append(s.items, item)
}

func (s *Stack[T]) Pop(item *T) bool {
	if item == nil || len(s.items) == 0 {
		return false
	}
	*item = s.items[len(s.items)-1]
	s.items = s.items[:len(s.items)-1]
	return true
}
