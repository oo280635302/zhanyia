package program

import "fmt"

func Ingress() {
	fmt.Println(getNumberOfBacklogOrders([][]int{{26, 7, 0}, {16, 1, 1}, {14, 20, 0}, {23, 15, 1}, {24, 26, 0}, {19, 4, 1}, {1, 1, 0}}))
}
