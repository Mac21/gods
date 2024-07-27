package gods

type Stack[T any] []T

func (s *Stack[T]) Push(v T) {
	*s = append(*s, v)
}

func (s *Stack[T]) Pop() (T, bool) {
	if len(*s) < 1 {
        var zv T
		return zv, true
	}

    res := (*s)[len(*s)-1]
	*s = (*s)[:len(*s)-1]

	return res, false
}

func (s Stack[T]) Len() int {
	return len(s)
}
