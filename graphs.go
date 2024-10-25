package graphs

type Edge[T any] interface {
	From() Node[T]
	To() Node[T]
}

/* Common node interface for the graph */
type Node[T any] interface {
	OutgoingEdges() []Edge[T]
	IncomingEdges() []Edge[T]
	Neighbors() []Node[T]
	Value() T
}

type Graph[T any] interface {
	Size() int
	Nodes() []Node[T]
	/* TODO: add common graph metrics like girth, etc etc */
}

type WeightedEdge struct {
	Weight int
}
