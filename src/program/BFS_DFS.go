package program

// 广度优先搜索 + 深度优先搜索

// 判断二分图---------------------------------------------------------------------------------------------------------
// 思路: 广度优先搜索 用颜色来表示 当前节点与相邻节点染不同色 每次染色都检测周边是否染色 已染色的是否与当前需要染色的颜色不同错/相同对并染色	32ms/6.1mb
func IsBipartite(graph [][]int) bool {
	const (
		UNCOLORED, RED, GREEN = 0, 1, 2
	)

	n := len(graph)

	color := make([]int, n)
	for i := 0; i < n; i++ {
		if color[i] == UNCOLORED {
			queue := []int{}
			queue = append(queue, i)
			color[i] = RED

			for i := 0; i < len(queue); i++ {
				node := queue[i]
				cNei := RED
				if color[node] == RED {
					cNei = GREEN
				}
				for _, neighbor := range graph[node] {
					if color[neighbor] == UNCOLORED {
						queue = append(queue, neighbor)
						color[neighbor] = cNei
					} else if color[neighbor] != cNei {
						return false
					}
				}
			}
		}
	}
	return true
}
