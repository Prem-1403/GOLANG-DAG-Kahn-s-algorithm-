package dag

import (
	"fmt"
	"os"
)

// Print DAG information to console
func PrettyPrint(vertices map[int]Vertex, edges []Edge, indent int) {
	fmt.Println("------ VERTICES ------")
	for id, vertex := range vertices {
		fmt.Printf("%*sVertex %d: %v\n", indent, "", id, vertex.Data)
	}
	fmt.Println("------ EDGES ------")
	for _, edge := range edges {
		fmt.Printf("%*sEdge %d -> %d\n", indent, "", edge.Src, edge.Dest)
	}
}

// Check if file or directory exists
func PathExists(path string) bool {
	_, err := os.Stat(path)
	return err == nil || os.IsExist(err)
}
