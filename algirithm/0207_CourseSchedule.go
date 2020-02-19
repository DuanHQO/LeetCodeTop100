package algirithm

var markedCanFinish []bool
var onStack []bool
var adj [][]int
var hasCycleCanFinish bool

func CanFinish(numCourses int, prerequisites [][]int) bool {
	markedCanFinish = make([]bool, numCourses)
	onStack = make([]bool, numCourses)
	adj = make([][]int, numCourses)
	hasCycleCanFinish = false

	for _, pair := range prerequisites {
		v := pair[1]
		w := pair[0]
		adj[v] = append(adj[v], w)
	}

	for v := 0; v < numCourses; v++ {
		if !markedCanFinish[v] {
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
	markedCanFinish[v] = true
	onStack[v] = true
	for _, w := range adj[v] {
		if !markedCanFinish[w] {
			dfsCanFinish(w)
		} else if onStack[w] {
			hasCycleCanFinish = true
			return
		}
	}
	onStack[v] = false
}
