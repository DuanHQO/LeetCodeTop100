package algirithm

import (
	"leetcodetop100/kit"
)

func calcEquation(equations [][]string, values []float64, queries [][]string) []float64 {
	m := make(map[string]map[string]float64)

	for i, e := range equations {
		a, b := e[0], e[1]
		v := values[i]

		if _, ok := m[a]; !ok {
			m[a] = make(map[string]float64)
		}
		m[a][b] = 1.0 / v

		if _, ok := m[b]; !ok {
			m[b] = make(map[string]float64)
		}
		m[b][a] = v
	}

	res := make([]float64, len(queries))
	for i, q := range queries {
		res[i] = bfs(m, q[0], q[1])
	}

	return res
}

type entry struct {
	s string
	f float64
}

//广度优先搜索
func bfs(m map[string]map[string]float64, a string, b string) float64 {
	_, ok := m[a]
	if !ok {
		return -1.0
	}

	_, ok = m[b]
	if !ok {
		return -1
	}

	if a == b {
		return 1.0
	}

	isVisited := make(map[string]bool)
	queue := []entry{{a, 1.0}}

	for len(queue) > 0 {
		e := queue[0]
		queue = queue[1:]
		if e.s == b {
			return 1.0 / e.f
		}

		if isVisited[e.s] {
			continue
		}

		isVisited[e.s] = true

		for k, v := range m[e.s] {
			queue = append(queue, entry{k, v * e.f})
		}
	}

	return -1
}

func calcEquationA(equations [][]int, values []float64, queries [][]int) []float64 {
	dic := make(map[int]int)

	for _, equation := range equations {
		_, exist := dic[equation[0]]
		if !exist {
			dic[equation[0]] = 1
		}
	}

	digraph := kit.NewEdgeWeightedDigraph(len(dic))
	for i, equation := range equations {
		edge := kit.NewDirectedEdge(equation[0], equation[1], values[i])
		digraph.AddEdge(edge)
	}

}
