package program

import "fmt"

func Ingress() {
	c := platesBetweenCandles("**|**|***|", [][]int{{2, 5}, {5, 9}})

	fmt.Println(c)
}
