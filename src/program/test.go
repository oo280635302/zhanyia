package program

import "fmt"

func Ingress() {
	c := ConstructorFreqStack()
	c.Push(7)
	c.Push(7)
	c.Push(7)
	c.Push(1)
	fmt.Println(c.Pop())
	fmt.Println(c.Pop())
	fmt.Println(c.Pop())
	fmt.Println(c.Pop())
}
