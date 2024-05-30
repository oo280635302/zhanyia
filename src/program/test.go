package program

import "fmt"

func Ingress() {
	fmt.Println(flipChess([]string{".X.", ".O.", "XO."}))
}
