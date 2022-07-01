package testdata

type Node[T any] struct {
}

func (n *Node[T]) String() string { // want "cognitive complexity 1 of func \\(\\*Node\\[T\\]\\)\\.String is high \\(> 0\\)"
	if n != nil { // +1
		return "Node"
	}

	return ""
} // total complexity = 1

type Pair[K any, V any] struct {
	Key   K
	Value V
}

func (p *Pair[K, V]) String() string { // want "cognitive complexity 1 of func \\(\\*Pair\\[K, V\\]\\)\\.String is high \\(> 0\\)"
	if p != nil { // +1
		return "Pair"
	}

	return ""
} // total complexity = 1

type Triple[K any, V any, T any] struct {
}

func (t *Triple[K, V, T]) String() string { // want "cognitive complexity 1 of func \\(\\*Triple\\[K, V, T\\]\\)\\.String is high \\(> 0\\)"
	if t != nil { // +1 `
		return "Triple"
	}

	return ""
} // total complexity = 1
