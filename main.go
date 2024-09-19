package main

import (
	"blockdag-explorer/dag"
	"fmt"
	"sort"
)

func main() {
	// Initialize DAG with genesis block
	dagObj := dag.NewBlockDAG()
	genesisBlock := dag.NewBlock(0, nil)
	dagObj.AddBlock(genesisBlock)

	fmt.Println("\nSelect an option:")
	fmt.Println("1) Add the Block to the DAG")
	fmt.Println("2) Add at Specific Position")
	fmt.Println("3) Exit")

	for {
		var choice int
		fmt.Print("Enter your choice (1, 2, or 3): ")
		_, err := fmt.Scan(&choice)
		if err != nil || choice == 3 || choice == -1 {
			break
		}

		switch choice {
		case 1:
			if !addBlockToDAG(dagObj) {
				return // Exit the program
			}
		case 2:
			if !addBlockAtSpecificPosition(dagObj) {
				return // Exit the program
			}
		default:
			fmt.Println("Invalid choice. Please select 1, 2, or 3.")
		}
	}

	// Visualize DAG
	dag.VisualizeDAG(dagObj, "assets/block_dag.html")
}

func addBlockToDAG(dagObj *dag.BlockDAG) bool {
	var id int

	fmt.Print("Enter block ID (or -1 to exit): ")
	_, err := fmt.Scan(&id)
	if err != nil || id == -1 {
		return false // Exit the console
	}

	// Add the new block (will automatically connect in linear chain)
	block := dag.NewBlock(id, nil)
	err = dagObj.AddBlock(block)
	if err != nil {
		fmt.Println(err)
		return true // Continue to allow further inputs
	}

	// Sort and connect blocks in order
	connectBlocksSequentially(dagObj)

	return true
}

func addBlockAtSpecificPosition(dagObj *dag.BlockDAG) bool {
	var id, parentID int

	fmt.Print("Enter block ID (or -1 to exit): ")
	_, err := fmt.Scan(&id)
	if err != nil || id == -1 {
		return false // Exit the function but not the program
	}

	fmt.Print("Enter parent block ID to connect to (or -1 to exit): ")
	_, err = fmt.Scan(&parentID)
	if err != nil || parentID == -1 {
		return false // Exit the function but not the program
	}

	// Check if the parent block exists in the DAG
	parentBlock, exists := dagObj.Blocks[parentID]
	if !exists {
		fmt.Println("Invalid parent block ID.")
		return true // Allow user to retry
	}

	// Add the new block at the specified position
	block := dag.NewBlock(id, []*dag.Block{parentBlock})
	err = dagObj.AddBlock(block)
	if err != nil {
		fmt.Println(err)
		return true // Allow user to retry
	}

	return true
}

func connectBlocksSequentially(dagObj *dag.BlockDAG) {
	// Extract all block IDs, sort them, and connect sequentially to form a chain
	var blockIDs []int
	for id := range dagObj.Blocks {
		blockIDs = append(blockIDs, id)
	}
	sort.Ints(blockIDs)

	// Connect blocks in sorted order to the genesis block (0)
	if len(blockIDs) == 0 {
		return // No blocks to connect
	}

	previousBlock := dagObj.Blocks[blockIDs[0]] // Start from genesis block
	for _, id := range blockIDs[1:] {           // Skip genesis block
		block := dagObj.Blocks[id]
		block.Parent = []*dag.Block{previousBlock}
		previousBlock = block
	}
}
