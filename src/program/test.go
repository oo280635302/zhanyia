package program

import (
	"fmt"
	"time"
)

func Ingress() {
	s := time.Now().UnixNano()

	fmt.Println(maxSumTwoNoOverlap([]int{1, 2, 3, 4, 5, 6, 7, 8}, 1, 2))

	fmt.Println("耗时：", (time.Now().UnixNano()-s)/1e6)

}
