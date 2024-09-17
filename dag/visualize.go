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
			SymbolSize: 20,
			// Optionally, use 'arrow' or another symbol for target nodes if available
		}
		nodes = append(nodes, node)
	}

	// For every edge, create a corresponding link (without arrows for now)
	for _, edge := range edges {
		source := fmt.Sprintf("Vertex %d", edge.Src)
		target := fmt.Sprintf("Vertex %d", edge.Dest)
		link := opts.GraphLink{
			Source: source,
			Target: target,
		}
		links = append(links, link)
	}

	// Debug print to check the generated nodes and links
	fmt.Println("Nodes:", nodes)
	fmt.Println("Links:", links)
	tip := true
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

	f, err := os.Create(filePath)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	graph.Render(f)
}
