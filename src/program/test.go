package program

import "fmt"

func Ingress() {
	c := luckyNumbers([][]int{{3, 7, 8}, {9, 11, 13}, {5, 16, 17}})

	fmt.Println(c)
}
