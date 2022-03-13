package program

import "fmt"

func Ingress() {
	c := validUtf8([]int{250, 145, 145, 145, 145})

	fmt.Println(c)
}
