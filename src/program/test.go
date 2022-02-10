package program

import "fmt"

func Ingress() {
	c := countKDifference([]int{1, 2, 2, 1}, 1)

	fmt.Println(c)
}
