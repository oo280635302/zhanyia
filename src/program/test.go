package program

import "fmt"

func Ingress() {
	a := containsNearbyDuplicate([]int{11, 12, 13, 14, 15}, 2)
	fmt.Println(a)
}
