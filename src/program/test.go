package program

import "fmt"

func Ingress() {
	c := getFolderNames([]string{"kaido", "kaido(1)", "kaido", "kaido(1)(1)", "kaido(1)"})

	fmt.Println(c)
}
