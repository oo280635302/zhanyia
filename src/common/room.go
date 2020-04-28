package common

import (
	"sync"
	"time"
	pb "zhanyia/src/proto"
)

// 单个房间
type OneRoom struct {
	// 房间锁
	lock sync.Mutex
	// 房间信息
	RoomInfo *pb.RoomInfo
	// 定时器协程
	timer chan bool
}

// 房间驱动
func (r *OneRoom) Drive(roomDriveFunc func(roomInfo *pb.RoomInfo) *pb.RoomInfo) {
	r.lock.Lock()
	defer r.lock.Unlock()
	// 关闭上个定时器协程
	r.timer <- true

	// 运行driver
	r.RoomInfo = roomDriveFunc(r.RoomInfo)

	GoTimer(time.Millisecond*5000, func() bool {
		r.Drive(roomDriveFunc)
		return false
	})
}

// 保存房间
func (r *OneRoom) save() {

}
