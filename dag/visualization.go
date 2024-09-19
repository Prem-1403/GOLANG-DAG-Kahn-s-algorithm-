package dag

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/go-echarts/go-echarts/v2/charts"
	"github.com/go-echarts/go-echarts/v2/opts"
)

// VisualizeDAG creates a visual representation of the DAG and saves it as an HTML file
func VisualizeDAG(dag *BlockDAG, filePath string) {
	// Ensure the directory exists
	dir := filepath.Dir(filePath)
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		os.MkdirAll(dir, os.ModePerm)
	}

	graph := charts.NewGraph()

	var nodes []opts.GraphNode
	var links []opts.GraphLink

	// Create nodes for each block using the Block ID and Hash
	for _, block := range dag.Blocks {
		nodeName := fmt.Sprintf("Block %d\nHash: %s", block.ID, block.Hash[:8]) // Display first 8 chars of the hash for readability
		node := opts.GraphNode{
			Name:       nodeName,
			SymbolSize: 20, // Increase the size of the nodes
		}
		nodes = append(nodes, node)

		// Create links between parent and child blocks using Block ID
		for _, parentBlock := range block.Parent {
			parentNodeName := fmt.Sprintf("Block %d\nHash: %s", parentBlock.ID, parentBlock.Hash[:8])
			link := opts.GraphLink{
				Source: parentNodeName,
				Target: nodeName,
				LineStyle: &opts.LineStyle{
					Curveness: 0.2, // Adjust the curvature for a vertical/horizontal feel
				},
			}
			links = append(links, link)
		}
	}

	// Tooltip enabled
	tip := true

	// Set global options
	graph.SetGlobalOptions(
		charts.WithTitleOpts(opts.Title{
			Title: "Block DAG Visualization",
		}),
		charts.WithTooltipOpts(opts.Tooltip{
			Show: &tip,
		}),
		charts.WithInitializationOpts(opts.Initialization{
			Width:  "1200px",
			Height: "800px",
		}),
	)

	graph.AddSeries("BlockDAG", nodes, links).
		SetSeriesOptions(
			charts.WithGraphChartOpts(opts.GraphChart{
				Force: &opts.GraphForce{
					Repulsion: 1600,
				},
				EdgeSymbol:     []string{"none", "arrow"}, // Show arrows on edges
				EdgeSymbolSize: 10,
			}),
		)

	// Save graph to an HTML file
	f, err := os.Create(filePath)
	if err != nil {
		panic(err)
	}
	defer f.Close()
	graph.Render(f)
}
