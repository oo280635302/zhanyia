package program

import "fmt"

func Ingress() {
	x := ConstructorThrone("king")
	x.Birth("king", "1")
	x.Birth("king", "2")
	x.Birth("king", "3")
	x.Birth("king", "4")
	x.Birth("2", "5")
	fmt.Println(x.GetInheritanceOrder())
}
