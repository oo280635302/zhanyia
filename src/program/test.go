package program

import "fmt"

func Ingress() {
	fmt.Println(applyOperations([]int{1, 2, 2, 1, 1, 0}))
}
