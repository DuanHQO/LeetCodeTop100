package kit

type EdgeWeightedDigraph struct {
	V   int
	E   int
	adj [][]DirectedEdge
}

func NewEdgeWeightedDigraph(v int) *EdgeWeightedDigraph {
	digraph := &EdgeWeightedDigraph{
		V:   v,
		E:   0,
		adj: make([][]DirectedEdge, v),
	}
	return digraph
}

func (this *EdgeWeightedDigraph) AddEdge(e DirectedEdge) {
	this.adj[e.From()] = append(this.adj[e.From()], e)
	this.E++
}

func (this *EdgeWeightedDigraph) Edges() []DirectedEdge {
	var bag []DirectedEdge
	for v := 0; v < this.V; v++ {
		for _, edge := range this.adj[v] {
			bag = append(bag, edge)
		}
	}
	return bag
}

func (this *EdgeWeightedDigraph) HasCycle() bool {
	return false
}
