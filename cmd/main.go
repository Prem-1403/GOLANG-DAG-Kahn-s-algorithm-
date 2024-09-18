package main

import (
	"blockdag/dag"
	"fmt"
)

func main() {
	vertices, edges := dag.LoadData("C:\\Users\\Prem\\Desktop\\DAG\\data\\sample_Dag.json")
	// Check for cycles before building the BlockDAG
	if dag.HasCycle(vertices, edges) {
		fmt.Println("Error: A cycle was detected in the DAG. This input does not form a valid DAG.")
		return
	}

	// Build BlockDAG
	signatures, err := dag.BuildBlockDAG(vertices, edges)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	// Print block signatures
	/*fmt.Println("Block Signatures:")
	for id, sig := range signatures {
		if id != -1 {
			fmt.Printf("Block %d: %s\n", id, sig["data_hash"])
		}
	}*/

	// Visualize DAG
	dag.VisualizeDAG(vertices, edges, "./blockdag.html")

	// Accessing signature with int key (-1)
	fmt.Println("BlockDAG signature:", signatures[-1])
}
