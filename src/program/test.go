package program

import "fmt"

func Ingress() {
	c := findClosest([]string{"I", "am", "a", "student", "from", "a", "university", "in", "a", "city"}, "a", "student")
	fmt.Println(c)
}
