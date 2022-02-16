package main

import "fmt"

// given a directed graph, design an algorithm to find out wheter there is a route between two nodes
// a directed graph is a graph

type Node struct {
	name string
}

type Graph struct {
	nodes []Node
	edges map[Node][]*Node
}

func (g *Graph) AddNode(n Node) Node {
	g.nodes = append(g.nodes, n)

	return n
}

func (g *Graph) AddEdge(n1 *Node, n2 *Node) {

	if g.edges == nil {
		g.edges = make(map[Node][]*Node)
	}
	g.edges[*n1] = append(g.edges[*n1], n2)
}

func (g *Graph) Print() {
	for i := 0; i < len(g.nodes); i++ {
		node := g.nodes[i]
		fmt.Printf("%s ->", node.name)
		edges := g.edges[node]
		for j := 0; j < len(edges); j++ {
			fmt.Printf(" %s ", edges[j].name)
		}
		fmt.Printf("\n")
	}
}

func (g *Graph) hasRoute(n1 *Node, n2 *Node) bool {
	edges := g.edges[*n1]
	for i := 0; i < len(edges); i++ {
		if edges[i] == n2 || g.hasRoute(edges[i], n2) {
			return true
		}
	}
	return false
}

func main() {
	myGraph := Graph{}
	A := myGraph.AddNode(Node{"A"})
	B := myGraph.AddNode(Node{"B"})
	C := myGraph.AddNode(Node{"C"})
	D := myGraph.AddNode(Node{"D"})
	E := myGraph.AddNode(Node{"E"})

	myGraph.AddEdge(&A, &B)
	myGraph.AddEdge(&A, &C)
	myGraph.AddEdge(&B, &D)
	myGraph.AddEdge(&C, &E)

	myGraph.Print()

	fmt.Printf("Testing if there's a route between %s and %s: %v \n", A.name, D.name, myGraph.hasRoute(&A, &D))
	fmt.Printf("Testing if there's a route between %s and %s: %v \n", B.name, E.name, myGraph.hasRoute(&B, &E))

}
