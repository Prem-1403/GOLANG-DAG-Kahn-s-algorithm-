package main

import (
	"blockdag/dag"
	"fmt"
)

func main() {
	vertices, edges := dag.LoadData("C:/Users/Prem/Desktop/blockdag/data/sample_dag.json")

	// Check if data is loaded correctly
	fmt.Println("Loaded Vertices:", vertices)
	fmt.Println("Loaded Edges:", edges)

	// Build BlockDAG
	signatures, err := dag.BuildBlockDAG(vertices, edges)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	// Print block signatures
	fmt.Println("Block Signatures:")
	for id, sig := range signatures {
		if id != -1 {
			fmt.Printf("Block %d: %s\n", id, sig["data_hash"])
		}
	}

	// Visualize DAG
	dag.VisualizeDAG(vertices, edges, "./blockdag.html")

	// Accessing signature with int key (-1)
	fmt.Println("BlockDAG signature:", signatures[-1])
}
