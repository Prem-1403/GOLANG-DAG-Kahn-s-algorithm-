package dag

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/go-echarts/go-echarts/v2/charts"
	"github.com/go-echarts/go-echarts/v2/opts"
)

func VisualizeDAG(vertices map[int]Vertex, edges []Edge, filePath string) {
	dir := filepath.Dir(filePath)
	if !PathExists(dir) {
		os.MkdirAll(dir, os.ModePerm)
	}

	graph := charts.NewGraph()

	var nodes []opts.GraphNode
	var links []opts.GraphLink

	// For every vertex, create a corresponding node
	for _, vertex := range vertices {
		nodeName := fmt.Sprintf("Vertex %d", vertex.ID)
		node := opts.GraphNode{
			Name:       nodeName,
			SymbolSize: 5,
		}
		if vertex.ID == 0 {
			node.ItemStyle = &opts.ItemStyle{
				Color: "#FF0000", // Genesis block in red
			}
			node.SymbolSize = 10 // Make the genesis block larger
		}
		nodes = append(nodes, node)
	}

	// For every edge, create a corresponding link (with arrows)
	for _, edge := range edges {
		source := fmt.Sprintf("Vertex %d", edge.Src)
		target := fmt.Sprintf("Vertex %d", edge.Dest)
		link := opts.GraphLink{
			Source: source,
			Target: target,
		}
		links = append(links, link)
	}

	tip := true
	graph.SetGlobalOptions(
		charts.WithTitleOpts(opts.Title{
			Title: "Block DAG Visualization",
		}),
		charts.WithTooltipOpts(opts.Tooltip{
			Show: &tip,
		}),
		charts.WithInitializationOpts(opts.Initialization{
			Width:  "800%", // Use percentage for width
			Height: "800%", // Use percentage for height
		}),
	)
	graph.AddSeries("BlockDAG", nodes, links).
		SetSeriesOptions(
			charts.WithGraphChartOpts(opts.GraphChart{
				Force: &opts.GraphForce{
					Repulsion: 100, // Increased repulsion for large datasets
				},
				EdgeSymbol:     []string{"none", "arrow"}, // Show arrows on edges
				EdgeSymbolSize: 5,
				Roam:           &tip, // Enable zoom and pan
			}),
		)

	f, err := os.Create(filePath)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	graph.Render(f)
}
