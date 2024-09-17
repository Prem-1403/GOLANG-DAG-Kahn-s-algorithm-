package dag

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
)

type Vertex struct {
	ID   int               `json:"id"`
	Data map[string]string `json:"data"`
}

type Edge struct {
	Src  int `json:"src"`
	Dest int `json:"dest"`
}

func LoadVertices(filepath string) map[int]Vertex {
	file, _ := os.Open(filepath)
	defer file.Close()

	var vertices map[int]Vertex
	json.NewDecoder(file).Decode(&vertices)
	return vertices
}

func LoadEdges(filepath string) []Edge {
	file, _ := os.Open(filepath)
	defer file.Close()

	var edges []Edge
	json.NewDecoder(file).Decode(&edges)
	return edges
}

func BuildBlockDAG(vertices map[int]Vertex, edges []Edge) (map[int]map[string]string, error) {
	workingset := make(map[int]map[string]string)
	queue := make([]int, 0)

	// Initialize workingset and set in-degree of vertices
	inDegree := make(map[int]int)
	for id := range vertices {
		workingset[id] = make(map[string]string)
		inDegree[id] = 0
	}

	for _, edge := range edges {
		inDegree[edge.Dest]++
	}

	// Find vertices with zero in-degree
	for id, degree := range inDegree {
		if degree == 0 {
			queue = append(queue, id)
		}
	}

	// Process the DAG
	var visited []int
	for len(queue) > 0 {
		var vID int
		vID, queue = queue[len(queue)-1], queue[:len(queue)-1]

		workingset[vID]["data_hash"] = generateHash(vertices[vID].Data)

		// Process the edges
		for _, edge := range edges {
			if edge.Src == vID {
				workingset[edge.Dest]["parent_hashes"] = appendHash(workingset[edge.Src]["data_hash"])
				inDegree[edge.Dest]--

				if inDegree[edge.Dest] == 0 {
					queue = append(queue, edge.Dest)
				}
			}
		}

		visited = append(visited, vID)
	}

	if len(visited) != len(vertices) {
		return nil, errors.New("cycle detected in DAG")
	}

	// Build final signatures
	leaves := make([]map[string]string, 0)
	for id := range vertices {
		if isLeaf(vertices, edges, id) {
			leaves = append(leaves, workingset[id])
		}
	}

	finalSignature := generateGraphSignature(leaves)
	workingset[-1] = map[string]string{"hash": finalSignature}
	return workingset, nil
}
func LoadData(filepath string) (map[int]Vertex, []Edge) {
	file, err := os.Open(filepath)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return nil, nil
	}
	defer file.Close()

	var data struct {
		Vertices map[int]Vertex `json:"vertices"`
		Edges    []Edge         `json:"edges"`
	}

	if err := json.NewDecoder(file).Decode(&data); err != nil {
		fmt.Println("Error decoding JSON:", err)
		return nil, nil
	}

	return data.Vertices, data.Edges
}
func HasCycle(vertices map[int]Vertex, edges []Edge) bool {
	visited := make(map[int]bool)
	recStack := make(map[int]bool)

	// Helper function for DFS cycle detection
	var dfs func(v int) bool
	dfs = func(v int) bool {
		if recStack[v] {
			return true
		}
		if visited[v] {
			return false
		}

		visited[v] = true
		recStack[v] = true

		for _, edge := range edges {
			if edge.Src == v { // Capitalized field
				if dfs(edge.Dest) { // Capitalized field
					return true
				}
			}
		}

		recStack[v] = false
		return false
	}

	// Check each vertex for cycle
	for vertex := range vertices {
		if dfs(vertex) {
			return true
		}
	}

	return false
}

func isLeaf(vertices map[int]Vertex, edges []Edge, id int) bool {
	if _, exists := vertices[id]; !exists {
		return false
	}
	for _, edge := range edges {
		if edge.Src == id {
			return false
		}
	}
	return true
}
