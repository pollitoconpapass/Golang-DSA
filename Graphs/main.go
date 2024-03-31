package main

import "fmt"

// === GRAPH DECLARATION ===
type Graph struct {
	vertices []*Vertex
}

// === VERTEX DECLARATION ===
type Vertex struct {
	value    int
	adjacent []*Vertex // -> the vertexes which are connected to the current vertex
}

// === ADD VERTEX ===
func (g *Graph) AddVertex(value int) {
	if doesItContains(g.vertices, value) { // -> does the validation first
		err := fmt.Errorf("Vertex %v already exists", value)
		fmt.Println(err.Error()) // -> print the error handling
		return
	}

	g.vertices = append(g.vertices, &Vertex{value: value}) // -> append the new vertex to the list
}

// === ADD EDGE ===
func (g *Graph) addEdge(from, to int) {
	// Get the vertex
	fromVertex := g.getVertex(from) // -> (from: origin)
	toVertex := g.getVertex(to)     // ->  (to: vertex we want to connect

	// Check errors
	if fromVertex == nil || toVertex == nil { // -> if the origin vertex doesn't exist
		err := fmt.Errorf("Invalid edge (%v -> %v)", from, to)
		fmt.Println(err.Error())
		return
	} else if doesItContains(fromVertex.adjacent, to) { // -> if the connection already exists
		err := fmt.Errorf("Already existing edge (%v -> %v)", from, to)
		fmt.Println(err.Error())
	}

	fromVertex.adjacent = append(fromVertex.adjacent, toVertex) // Add the edge to the vertex
}

// === GET VERTEX ===
func (g *Graph) getVertex(value int) *Vertex { // -> find the vertex given the value (if it exists)
	for i, v := range g.vertices { // -> loop over all the vertices
		if v.value == value { // -> if we find the vertice value we want
			return g.vertices[i] // -> return the vertex
		}
	}
	return nil // -> in case it didn't find the vertex
}

// === IF CONTAINS ===
func doesItContains(vertexList []*Vertex, value int) bool { // -> to avoid adding a value that already exists
	for _, v := range vertexList { // -> iterate over the list
		if value == v.value { // -> find if the value is already in the list
			return true
		}
	}
	return false
}

// === PRINT GRAPH ===
func (g *Graph) Print() {
	for _, v := range g.vertices {
		fmt.Printf("\nVertex %v", v.value)
		for _, v := range v.adjacent {
			fmt.Printf(" -> %v ", v.value)
		}
	}
	fmt.Printf("\n")
}

func main() {
	test := Graph{}

	for i := 0; i < 5; i++ {
		test.AddVertex(i)
	}

	test.AddVertex(0) // |
	test.AddVertex(0) // |-> this will print an error because the vertex already exists
	test.AddVertex(30)

	test.addEdge(1, 2) // -> (origin vertex, the other vertex we want to connect) ... In here, we want to connect vertex 1 to vertex 2 (1 as the origin vertex)
	test.addEdge(4, 1)
	test.addEdge(2, 3)

	test.addEdge(1, 2) // -> this will print an error because the edge connection already exists
	test.addEdge(6, 1) // -> this will print an error because the origin vertex doesn't exist

	test.Print()
}
