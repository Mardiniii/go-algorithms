package main

func main() {
	vertices := []string{"A", "B", "C", "D", "E"}

	edges := []edge{
		{from: "S", to: "A", cost: 4},
		{from: "S", to: "E", cost: 2},
		{from: "A", to: "D", cost: 3},
		{from: "A", to: "C", cost: 6},
		{from: "A", to: "B", cost: 5},
		{from: "B", to: "A", cost: 3},
		{from: "C", to: "B", cost: 1},
		{from: "D", to: "C", cost: 3},
		{from: "D", to: "A", cost: 1},
		{from: "E", to: "D", cost: 1},
	}

	result := NewResultTable(vertices)

	result.ApplyDijkstra(edges, result.origin.vertex)

	result.Print()
}
