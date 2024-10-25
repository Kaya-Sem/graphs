package graphs

type Edge interface {
	StartNode() Node
	EndNode() Node
}

/* Common node interface for the graph */
type Node interface {
	OutgoingEdges() []Edge
	IncomingEdges() []Edge
	CanocicalName() string
	TwinName() string
}

type Graph interface {
	size() int
}

type SimpleGraph struct {
	Size int
}

func (g SimpleGraph) New() SimpleGraph {
	graph := SimpleGraph{}

	return graph
}

type SimpleNode struct {
	twinFilePath string
	neighbors    []Node
	edges        []Edge
}

type SimpleEdge struct {
}

type WeightedEdge struct {
	Weight int
}
