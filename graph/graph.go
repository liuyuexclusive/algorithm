package graph

type Node struct {
	Val       int
	Neighbors []*Node
}

// 克隆图

// 给你无向 连通 图中一个节点的引用，请你返回该图的 深拷贝（克隆）。
// 图中的每个节点都包含它的值 val（int） 和其邻居的列表（list[Node]）。
func cloneGraph(node *Node) *Node {
	return cloneGraph2(node, make(map[int]*Node))
}

func cloneGraph2(node *Node, m map[int]*Node) *Node {
	if node == nil {
		return nil
	}
	if has, ok := m[node.Val]; ok {
		return has
	}
	res := &Node{Val: node.Val}
	m[node.Val] = res
	for i := 0; i < len(node.Neighbors); i++ {
		v := node.Neighbors[i]
		res.Neighbors = append(res.Neighbors, cloneGraph2(v, m))
	}
	return res
}

// 课程表

// 你这个学期必须选修 numCourse 门课程，记为 0 到 numCourse-1 。
// 在选修某些课程之前需要一些先修课程。 例如，想要学习课程 0 ，你需要先完成课程 1 ，我们用一个匹配来表示他们：[0,1]
// 给定课程总量以及它们的先决条件，请你判断是否可能完成所有课程的学习？
func canFinish(numCourses int, prerequisites [][]int) bool {
	parentNums := make([]int, numCourses)
	subs := make([][]int, numCourses)

	for _, v := range prerequisites {
		a, b := v[0], v[1]
		parentNums[a]++
		subs[b] = append(subs[b], a)
	}

	for numCourses > 0 {
		hasTarget := false
		var target int
		for k, v := range parentNums {
			if v == 0 {
				target = k
				hasTarget = true
				break
			}
		}
		if !hasTarget {
			return false
		}
		for _, v := range subs[target] {
			parentNums[v]--
		}
		parentNums[target] = -1
		numCourses--
	}
	return true
}
