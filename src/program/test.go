package program

import (
	"fmt"
)

func Ingress() {
	a := [][]int{
		{0, 0, 1},
		{0, 0, 0},
		{1, 1, 1},
	}
	fmt.Println(maxEqualRowsAfterFlips(a))
}
