package program

import "fmt"

func Ingress() {
	a := secondMinimum(2, [][]int{{1, 2}}, 1, 2)
	fmt.Println(a)
}
