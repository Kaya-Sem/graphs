package graphs

type SimpleGraph[T any] struct {
	size  int
	nodes []Node[T]
}

func (s SimpleGraph[T]) Nodes() []Node[T] {
	return s.nodes
}

func (s SimpleGraph[T]) Size() int {
	return s.size
}

func NewSimpleGraph[T any]() *SimpleGraph[T] {
	return &SimpleGraph[T]{size: 0}
}

func (g SimpleGraph[T]) New() SimpleGraph[T] {
	graph := SimpleGraph[T]{}

	return graph
}

/*
	Simple Nodes
*/

type SimpleNode[T any] struct {
	edges     []Edge[T]
	neighbors []Node[T]
	value     T
}

func (n SimpleNode[T]) Neighbors() []Node[T] {
	return n.neighbors
}

func (n SimpleNode[T]) Value() T {
	return n.value
}

func (n SimpleNode[T]) IncomingEdges() []Edge[T] {
	edges := []Edge[T]{}
	for _, e := range n.edges {
		if e.From() == Node[T](n) {
			edges = append(edges, e)
		}
	}

	return edges
}

func (n SimpleNode[T]) OutgoingEdges() []Edge[T] {
	edges := []Edge[T]{}
	for _, e := range n.edges {
		if e.From() == Node[T](n) {
			edges = append(edges, e)
		}
	}

	return edges
}

/*
	Simple Edges
*/

type SimpleEdge[T any] struct {
	from Node[T]
	to   Node[T]
}

func (e SimpleEdge[T]) From() Node[T] {
	return e.from
}

func (e SimpleEdge[T]) To() Node[T] {
	return e.to
}
