package program

import "fmt"

func Ingress() {
	c := ConstructorAllOne()
	c.Inc("hello")

	fmt.Println(c)
}
