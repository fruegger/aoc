package ds

type Queue[T any] struct {
	items []T
}

func (q *Queue[T]) Push(item T) {
	q.items = append(q.items, item)
}

func (q *Queue[T]) Pull(item *T) bool {
	if item == nil || len(q.items) == 0 {
		return false
	}
	*item = q.items[0]
	q.items = q.items[1:]
	return true
}

func (q *Queue[T]) Min(item *T, compare func(T, T) int) bool {
	if item == nil || len(q.items) == 0 {
		return false
	}
	*item = q.items[0]
	for i := 1; i < len(q.items); i++ {
		if compare(q.items[i], *item) < 0 {
			*item = q.items[i]
		}
	}
	return true
}
