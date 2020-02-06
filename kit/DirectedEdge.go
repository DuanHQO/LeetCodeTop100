package kit

type DirectedEdge struct {
	v      int
	w      int
	weight float64
}

func NewDirectedEdge(v, w int, weight float64) DirectedEdge {
	edge := DirectedEdge{
		v:      v,
		w:      w,
		weight: weight,
	}
	return edge
}

func (this *DirectedEdge) From() int {
	return this.v
}

func (this *DirectedEdge) To() int {
	return this.w
}
