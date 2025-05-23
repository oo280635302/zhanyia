package program

import "fmt"

func Ingress() {
	res := getValidPos([][]uint8{
		{0, 0, 0, 0, 0},
		{0, 1, 0, 0, 0},
		{0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0},
	}, 2, 2)
	fmt.Println(res)
}
