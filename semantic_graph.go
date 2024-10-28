package main

import (
	"fmt"
	"gopkg.in/yaml.v3"
	"log"
	"os"
	"path/filepath"
	"regexp"
)

func main() {
	graph := NewGraph()
	err := graph.Initialise()
	if err != nil {
		log.Fatalf("Error initialising: %v", err)
	}

	graph.AddNode(&Node{Type: RegularFile, Name: "node 1"})
	graph.AddNode(&Node{Type: RegularFile, Name: "node 2"})
	graph.AddNode(&Node{Type: RegularFile, Name: "node 3"})

	graph.PrintAdjacencyList()

}

const (
	twinFolder       = "/home/kaya-sem/.config/semantic_graph"
	twinRegexPattern = `^twin_.*\.yaml`
)

type Graph struct {
	Size int
	Root Node
}

// Constants for UNIX file types
const (
	RegularFile = 1
	Directory
	Symlink
	CharDevice
	BlockDevice
	FIFO
	Socket
	Internal = 0
)

func (g *Graph) Nodes() []Node {
	visited := make(map[*Node]bool)
	var nodes []Node

	var dfs func(n *Node)
	dfs = func(n *Node) {
		if visited[n] {
			return
		}
		visited[n] = true
		nodes = append(nodes, *n)
		for _, neighbor := range n.Neighbors() {
			dfs(neighbor)
		}
	}

	// Start DFS from the root node
	dfs(&g.Root)

	return nodes
}

func NewGraph() *Graph {
	return &Graph{
		Size: 1,
		Root: Node{Name: "Root", Type: Internal}}
}

func (g *Graph) AddNode(n *Node) {
	g.Root.ConnectTo(n)
}

func (g Graph) Initialise() error {
	regexPattern, err := regexp.Compile(twinRegexPattern)
	if err != nil {
		log.Fatalf("Failed to compile regex: %v", err)
	}

	// Walk through the directory to find matching files.
	err = filepath.Walk(twinFolder, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		// Check if the file matches the regex and is a regular file.
		if !info.IsDir() && regexPattern.MatchString(info.Name()) {
			fmt.Printf("Processing file: %s\n", info.Name())

			// Open and read the file.
			data, err := os.ReadFile(path)
			if err != nil {
				return fmt.Errorf("failed to read file %s: %w", path, err)
			}

			var parsedData map[string]interface{}

			if err := yaml.Unmarshal(data, &parsedData); err != nil {
				return fmt.Errorf("failed to parse YAML in file %s: %w", path, err)
			}

			// Do something with the parsed YAML data.
			// For example, just print it for now.
			fmt.Printf("Parsed YAML data for %s: %v\n", info.Name(), parsedData)
		}
		return nil
	})

	return nil

}

type DirectedEdge struct {
	from *Node
	to   *Node
}

type Node struct {
	Incoming []DirectedEdge
	Outgoing []DirectedEdge
	Value    map[string]interface{}
	Type     int
	Name     string
}

func (n *Node) Neighbors() []*Node {
	neighbors := []*Node{}

	for _, edge := range n.Incoming {
		neighbors = append(neighbors, edge.from)

	}

	for _, edge := range n.Outgoing {
		neighbors = append(neighbors, edge.to)

	}

	return neighbors
}

func (n1 *Node) ConnectTo(n2 *Node) {
	edge := DirectedEdge{from: n1, to: n2}
	n1.Outgoing = append(n1.Outgoing, edge)
	n2.Incoming = append(n2.Incoming, edge)
}

func (g *Graph) GetAdjacencyList() [][]*Node {
	// This will hold the adjacency list for each node
	adjacencyList := [][]*Node{}

	// Traverse all nodes; assume `g.Nodes()` returns a slice of all nodes in the graph.
	nodes := g.Nodes()
	for _, node := range nodes {
		// Collect all neighbors for the current node
		neighbors := node.Neighbors()
		adjacencyList = append(adjacencyList, neighbors)
	}

	return adjacencyList
}

// Helper function to print the adjacency list
func (g *Graph) PrintAdjacencyList() {
	adjacencyList := g.GetAdjacencyList()

	for i, neighbors := range adjacencyList {
		fmt.Printf("Node %d: ", i)
		for _, neighbor := range neighbors {
			fmt.Printf("%v ", neighbor.Name) // Print the node value (customize as needed)
		}
		fmt.Println()
	}
}
