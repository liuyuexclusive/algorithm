package graph

type Node struct {
	Val       int
	Neighbors []*Node
}

//克隆图
func cloneGraph(node *Node) *Node {
	m := make(map[int]*Node)
	return myClone(node, m)
}

func myClone(node *Node, m map[int]*Node) *Node {
	if node == nil {
		return nil
	}

	res := &Node{Val: node.Val}

	m[res.Val] = res

	for _, v := range node.Neighbors {
		if v2, ok := m[v.Val]; !ok {
			res.Neighbors = append(res.Neighbors, myClone(v, m))
		} else {
			res.Neighbors = append(res.Neighbors, v2)
		}
	}
	return res
}

//是够可以完成课程（拓扑排序）
func canFinish(numCourses int, prerequisites [][]int) bool {
	outCome := make([][]int, numCourses)
	inCome := make([]int, numCourses)
	for i := 0; i < len(prerequisites); i++ {
		outCome[prerequisites[i][1]] = append(outCome[prerequisites[i][1]], prerequisites[i][0])
		inCome[prerequisites[i][0]]++
	}
	for i := 0; i < numCourses; i++ {
		isok := false
		for j := 0; j < len(inCome); j++ {
			if inCome[j] == 0 {
				for _, v := range outCome[j] {
					inCome[v]--
				}
				isok = true
				inCome[j] = -1
				break
			}
		}
		if !isok {
			return false
		}
	}

	return true
}
