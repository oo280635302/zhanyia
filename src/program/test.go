package program

import (
	"fmt"
)

type CS struct {
	Name int64 `json:"name,default=123"`
}

func Ingress() {
	res := findTheDifference("abcd", "abecd")
	fmt.Println("res:", res)
}
