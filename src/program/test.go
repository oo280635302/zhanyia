package program

import "fmt"

func Ingress() {
	c := longestWord([]string{"a", "banana", "app", "appl", "ap", "apply", "apple"})
	fmt.Println(c)
}
