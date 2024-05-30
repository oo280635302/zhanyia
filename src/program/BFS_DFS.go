package program

import "fmt"

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

// 飞地的数量
// 思路：广度优先，找到边界上的所有为1的点位然后扩散，找到所有的1 - 边界扩散的1 = 飞的数量
func numEnclaves(grid [][]int) (ans int) {
	type pair struct{ x, y int }
	var dirs = []pair{{-1, 0}, {1, 0}, {0, -1}, {0, 1}} // 4个方位

	m, n := len(grid), len(grid[0])
	vis := make([][]bool, m)
	for i := range vis {
		vis[i] = make([]bool, n)
	}
	q := []pair{}
	for i, row := range grid {
		if row[0] == 1 { // 左边界的1
			vis[i][0] = true
			q = append(q, pair{i, 0})
		}
		if row[n-1] == 1 { // 右边界的1
			vis[i][n-1] = true
			q = append(q, pair{i, n - 1})
		}
	}
	for j := 1; j < n-1; j++ {
		if grid[0][j] == 1 { // 上边界的1
			vis[0][j] = true
			q = append(q, pair{0, j})
		}
		if grid[m-1][j] == 1 { // 下边界的1
			vis[m-1][j] = true
			q = append(q, pair{m - 1, j})
		}
	}
	for len(q) > 0 { // 往外扩张，找到所有相邻的1
		p := q[0]
		q = q[1:]
		for _, d := range dirs {
			if x, y := p.x+d.x, p.y+d.y; 0 <= x && x < m && 0 <= y && y < n && grid[x][y] == 1 && !vis[x][y] { // 只往有效的扩张，已被扩张的不重复扩张
				vis[x][y] = true
				q = append(q, pair{x, y}) // 扩张
			}
		}
	}
	for i := 1; i < m-1; i++ { // 找到所有的1
		for j := 1; j < n-1; j++ {
			if grid[i][j] == 1 && !vis[i][j] { // 当前1没被扩张说明就是飞地
				ans++
			}
		}
	}
	return
}

// 网络空闲的时刻
func networkBecomesIdle(edges [][]int, patience []int) (ans int) {
	n := len(patience)
	g := make([][]int, n)
	for _, e := range edges {
		x, y := e[0], e[1]
		g[x] = append(g[x], y)
		g[y] = append(g[y], x)
	}

	vis := make([]bool, n)
	vis[0] = true
	q := []int{0}
	for dist := 1; q != nil; dist++ {
		tmp := q
		q = nil
		for _, x := range tmp {
			for _, v := range g[x] {
				if vis[v] {
					continue
				}
				vis[v] = true
				q = append(q, v)
				ans = max(ans, (dist*2-1)/patience[v]*patience[v]+dist*2+1)
			}
		}
	}
	return
}

// 两数之和 IV - 输入 BST
func findTarget(root *TreeNode, k int) bool {
	m := make(map[int]bool)

	d := []*TreeNode{root}
	for len(d) != 0 {
		p := d[0]
		d = d[1:]

		m[p.Val] = true
		if p.Left != nil {
			d = append(d, p.Left)
		}
		if p.Right != nil {
			d = append(d, p.Right)
		}
	}
	fmt.Println(m)
	for key := range m {
		if m[k-key] && k-key != key {
			return true
		}
	}

	return false
}

// LCP 41. 黑白翻转棋
func flipChess(chessboard []string) (ans int) {
	m, n := len(chessboard), len(chessboard[0])
	bfs := func(i, j int) int {
		q := [][2]int{{i, j}}
		g := make([][]byte, m)
		for idx := range g {
			g[idx] = make([]byte, n)
			copy(g[idx], chessboard[idx])
		}
		g[i][j] = 'X'
		cnt := 0
		for len(q) > 0 {
			p := q[0]
			q = q[1:]
			i, j = p[0], p[1]
			for a := -1; a <= 1; a++ {
				for b := -1; b <= 1; b++ {
					if a == 0 && b == 0 {
						continue
					}
					x, y := i+a, j+b
					for x >= 0 && x < m && y >= 0 && y < n && g[x][y] == 'O' {
						x, y = x+a, y+b
					}
					if x >= 0 && x < m && y >= 0 && y < n && g[x][y] == 'X' {
						x -= a
						y -= b
						cnt += max(abs(x-i), abs(y-j))
						for x != i || y != j {
							g[x][y] = 'X'
							q = append(q, [2]int{x, y})
							x -= a
							y -= b
						}
					}
				}
			}
		}
		return cnt
	}
	for i, row := range chessboard {
		for j, c := range row {
			if c == '.' {
				ans = max(ans, bfs(i, j))
			}
		}
	}
	return
}
