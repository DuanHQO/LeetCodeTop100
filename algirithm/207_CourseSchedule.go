package algirithm

var marked []bool
var onStack []bool
var adj [][]int
var hasCycleCanFinish bool

func CanFinish(numCourses int, prerequisites [][]int) bool {
	marked = make([]bool, numCourses)
	onStack = make([]bool, numCourses)
	adj = make([][]int, numCourses)
	hasCycleCanFinish = false

	for _, pair := range prerequisites {
		v := pair[1]
		w := pair[0]
		adj[v] = append(adj[v], w)
	}

	for v := 0; v < numCourses; v++ {
		if !marked[v] {
			dfsCanFinish(v)
		}
	}

	if hasCycleCanFinish {
		return false
	} else {
		return true
	}
}

func dfsCanFinish(v int) {
	marked[v] = true
	onStack[v] = true
	for _, w := range adj[v] {
		if !marked[w] {
			dfsCanFinish(w)
		} else if onStack[w] {
			hasCycleCanFinish = true
			return
		}
	}
	onStack[v] = false
}
