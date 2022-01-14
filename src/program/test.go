package program

import "fmt"

func Ingress() {
	a := kSmallestPairs([]int{1, 3, 5, 7, 9, 11, 13}, []int{2, 4, 6, 8, 10, 12}, 3)
	fmt.Println(a)
}
