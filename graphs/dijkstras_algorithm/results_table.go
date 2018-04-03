package main

import (
	"fmt"
	"math"
)

// ResultTable store the results after apply Dijsktra algorithm
type ResultTable struct {
	origin *Row
	table  []*Row
}

// Exported methods

// Print in console the result for each vertex
func (t *ResultTable) Print() {
	for _, row := range t.table {
		fmt.Println(row)
	}
}

// ApplyDijkstra applies the dijkstra algorithm to find the shortest path
func (t *ResultTable) ApplyDijkstra(edges []edge, v string) {
	// Look for edges with from equals to the given vertex
	var vertexEdges []edge

	for _, edge := range edges {
		if edge.from == v {
			vertexEdges = append(vertexEdges, edge)
		}
	}

	for _, edge := range vertexEdges {
		from := edge.from
		fromRow := t.findRowByVertice(from)
		fromCost := fromRow.cost

		to := edge.to
		toRow := t.findRowByVertice(to)
		toCost := toRow.cost

		newCost := fromCost + edge.cost

		// Update the cost if we find less cost than the current one
		if newCost < toCost {
			toRow.SetCost(newCost)
		}
	}

	// Mark the vertex as visited for the next iterations
	t.findRowByVertice(v).SetAsVisited()

	// Look for next vertex
	t.nextVertex(edges)
}

// NewResultTable creates a new table with the origin vertex
// called S and the rest of the vertices from the graph. All of
// them will have cost 0 and visited equal to false for the beginning
func NewResultTable(vertices []string) ResultTable {
	o := &Row{vertex: "S", cost: 0, visited: false}

	resultTable := ResultTable{
		origin: o,
		table:  []*Row{o},
	}
	resultTable.buildVerticesRows(vertices)

	return resultTable
}

// Unexported methods

func (t *ResultTable) buildVerticesRows(vertices []string) {
	for _, v := range vertices {
		t.table = append(t.table, &Row{vertex: v, cost: math.MaxInt64, visited: false})
	}
}

func (t *ResultTable) findRowByVertice(v string) *Row {
	var r *Row

	for _, row := range t.table {
		if row.vertex == v {
			r = row
		}
	}

	return r
}

func (t *ResultTable) nextVertex(edges []edge) {
	var next *Row
	var unvisited []*Row

	// Find the unvisited vertices
	for _, row := range t.table {
		if row.visited == false {
			unvisited = append(unvisited, row)
		}
	}

	// Check which one from the unvisited vertices has
	// lower cost. It will be the next vertex to visit
	if len(unvisited) > 0 {
		lowestCost := math.MaxInt64
		for _, row := range unvisited {
			if row.cost < lowestCost {
				next = row
			}
		}
	}

	// Recursive call to apply the algorithm to the next vertex
	if next != nil {
		t.ApplyDijkstra(edges, next.vertex)
	}
}
