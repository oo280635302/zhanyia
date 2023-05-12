package program

import (
	"fmt"
	"time"
)

func Ingress() {
	s := time.Now().UnixNano()

	a := [][]byte{{'#', '#', '#', '#', '#', '#'},
		{'#', 'S', '#', '#', '#', '#'},
		{'#', 'T', '.', 'B', '.', '#'},
		{'#', '.', '#', '#', '.', '#'},
		{'#', '.', '.', '.', '.', '#'},
		{'#', '#', '#', '#', '#', '#'}}
	fmt.Println(minPushBox(a))

	fmt.Println("耗时：", (time.Now().UnixNano()-s)/1e6)

}
