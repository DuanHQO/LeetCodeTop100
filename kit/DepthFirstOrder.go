package kit

type DepthFirstOrder struct {
	marked        []bool
	Pre           *Queue
	Post          *Queue
	ReversePost   *Stack
	StartingPoint int
	EdgeTo        []DirectedEdge
}

func NewDepthFirstOrder(G *EdgeWeightedDigraph, s int) *DepthFirstOrder {
	tmp := &DepthFirstOrder{
		marked:        make([]bool, G.V),
		Pre:           NewQueue(),
		Post:          NewQueue(),
		ReversePost:   NewStack(),
		StartingPoint: s,
		EdgeTo:        make([]DirectedEdge, G.V),
	}

	for v := 0; v < G.V; v++ {
		if !tmp.marked[v] {
			tmp.DFS(G, v)
		}
	}

	tmp.StartingPoint = s

	return tmp
}

func (this *DepthFirstOrder) DFS(G *EdgeWeightedDigraph, v int) {
	this.Pre.Push(v)
	this.marked[v] = true
	for _, edge := range G.adj[v] {
		if this.marked[edge.To()] {
			this.EdgeTo[edge.To()] = edge
			this.DFS(G, edge.To())
		}
	}
	this.Post.Push(v)
	this.ReversePost.Push(v)
}

func (this *DepthFirstOrder) HasPathTo(v int) bool {
	return this.marked[v]
}

func (this *DepthFirstOrder) PathTo(v int) []int {
	if !this.HasPathTo(v) {
		return nil
	}
	var path []int
	for x := v; x != this.StartingPoint; x = this.EdgeTo[x].From() {
		tmp := []int{x}
		path = append(tmp, path...)
	}
	tmp := []int{this.StartingPoint}
	path = append(tmp, path...)
	return path
}
