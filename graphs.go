package graphs

type Edge[T any] interface {
	StartNode() Node[T]
	EndNode() Node[T]
}

/* Common node interface for the graph */
type Node[T any] interface {
	OutgoingEdges() []Edge[T]
	IncomingEdges() []Edge[T]
	Value() T
}

type Graph[T any] interface {
	Size() int
	Nodes() []Node[T]
	/* TODO: add common graph metrics like girth, etc etc */
}

type SimpleGraph[T any] struct {
	size  int
	nodes []Node[T]
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

type SimpleNode[T any] struct {
	neighbors []Node[T]
	edges     []Edge[T]
}

type SimpleEdge struct {
}

type WeightedEdge struct {
	Weight int
}
