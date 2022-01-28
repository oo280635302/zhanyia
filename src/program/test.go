package program

import "fmt"

func Ingress() {
	c := numberOfWeakCharacters([][]int{{2, 2}, {1, 3}, {2, 3}, {2, 1}, {3, 2}})

	fmt.Println(c)
}
