package program

import "fmt"

func Ingress() {
	c := numberOfGoodSubsets([]int{1, 1, 3, 5})

	fmt.Println(c)
}
