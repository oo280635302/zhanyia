package program

import "fmt"

func Ingress() {
	fmt.Println(discountPrices("1 2 $3 4 $5 $6 7 8$ $9 $10$", 100))
}
