package program

import "fmt"

func Ingress() {
	c := countMaxOrSubsets([]int{3, 1})
	fmt.Println(c)
}
