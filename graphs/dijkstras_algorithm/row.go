package main

// Row is used to model the result for each vertex and have
// tracking if it was visited before or not
type Row struct {
	vertex  string
	cost    int
	visited bool
}

// SetCost sets the given cost value
func (r *Row) SetCost(c int) {
	r.cost = c
}

// SetAsVisited sets visited to true
func (r *Row) SetAsVisited() {
	r.visited = true
}
