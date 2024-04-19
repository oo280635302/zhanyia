package program

import "fmt"

func Ingress() {
	fmt.Println(minSkips([]int{4, 4, 16, 20, 8, 8, 2, 10}, 5, 30))
}
