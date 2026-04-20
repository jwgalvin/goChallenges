package graphs_practice

// CanFinish returns true if all numCourses can be finished given the
// prerequisite pairs (cycle detection in a directed graph).
// Input: numCourses=2, prerequisites=[[1,0]]  →  true
// Input: numCourses=2, prerequisites=[[1,0],[0,1]]  →  false
func CanFinish(numCourses int, prerequisites [][]int) bool {
	// TODO: implement (DFS with 3-color marking or Kahn's topo sort)
	return false
}

// FindOrder returns a valid course order, or nil if impossible (topo sort).
// Input: numCourses=4, prerequisites=[[1,0],[2,0],[3,1],[3,2]]  →  [0,1,2,3]
func FindOrder(numCourses int, prerequisites [][]int) []int {
	// TODO: implement (Kahn's BFS topo sort)
	return nil
}
