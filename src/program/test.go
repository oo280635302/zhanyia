package program

import "fmt"

func Ingress() {
	//["a","a","a","a","a","b","b","b","b","b","b"]
	//	["23:20","11:09","23:30","23:02","15:28",]
	fmt.Println(permuteUnique([]int{1, 1, 1}))
}
