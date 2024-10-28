package graphs

type Graph[T any] struct {
	Size   int
	Primus Node[T]
}

func (s Graph[T]) Nodes() []Node[T] {
	/* TODO: full traversal */
	return nil
}

func NewGraph[T any]() *Graph[T] {
	return &Graph[T]{Size: 0}
	/* TODO: load primus from filesystem */
}

/* TODO: create function to initialize a graph from primus */

type DirectedEdge[T any] struct {
	from *Node[T]
	to   *Node[T]
}

type Node[T any] struct {
	edges []DirectedEdge[T]
	Value T
}

func (n Node[T]) Neighbors() []*Node[T] {
	neighbors := []*Node[T]{}
	for _, edge := range n.edges {
		neighbors = append(neighbors, edge.to)

	}

	return neighbors
}

func (n *Node[T]) IncomingEdges() []DirectedEdge[T] {
	edges := []DirectedEdge[T]{}
	for _, edge := range n.edges {
		if edge.to == n {
			edges = append(edges, edge)
		}
	}
	return edges
}

func (n *Node[T]) OutgoingEdges() []DirectedEdge[T] {
	edges := []DirectedEdge[T]{}
	for _, edge := range n.edges {
		if edge.from == n {
			edges = append(edges, edge)
		}
	}
	return edges
}
