package program

import "fmt"

func Ingress() {
	c := bestRotation([]int{2, 3, 1, 4, 0})

	fmt.Println(c)
}
