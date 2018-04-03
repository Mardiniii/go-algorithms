package main

// Edge struct is used to model the paths between each node
// using three attributes from and to vertices among the cost
// between them
type edge struct {
	from string
	to   string
	cost int
}
