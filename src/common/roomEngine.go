package common

import (
	"github.com/google/uuid"
	"sync"
	pb "zhanyia/src/proto"
)

// 所有房间
type RoomMap struct {
	// 整体锁
	lock sync.Mutex
	// 房间 - 房间号/房间
	room map[string]*OneRoom
	// 驱动回调
	RoomDriveFunc func(roomInfo *pb.RoomInfo) *pb.RoomInfo
}

// 启动驱动
func (rm *RoomMap) Run(roomDriveFunc func(roomInfo *pb.RoomInfo) *pb.RoomInfo) {
	rm.lock.Lock()
	rm.lock.Unlock()
	rm.room = make(map[string]*OneRoom)
	rm.RoomDriveFunc = roomDriveFunc
}

// 创建房间
func (rm *RoomMap) CreateRoom() {
	rm.lock.Lock()
	defer rm.lock.Unlock()
	room := &pb.RoomInfo{}
	room.RoomUuid = uuid.New().String()
	room.RoomPlayerInfo = make([]*pb.RoomPlayerInfo, 0)
	rm.room[room.RoomUuid].RoomInfo = room
	rm.room[room.RoomUuid].Drive(rm.RoomDriveFunc)
}

// 删除房间
func (rm *RoomMap) DeleteRoom(roomUuid string) bool {
	if roomUuid == "" {
		return false
	}
	rm.lock.Lock()
	defer rm.lock.Unlock()

	for k, _ := range rm.room {
		if k == roomUuid {
			delete(rm.room, k)
			return true
		}
	}
	return false
}

// 所有房间停止
func (rm *RoomMap) AllStop() {
	for _, v := range rm.room {
		v.timer <- true
	}
}

// 所有房间启动 - 需要停止后启动 - 如果没暂停启动 - 等着崩溃吧
func (rm *RoomMap) AllStart() {
	for _, v := range rm.room {
		v.Drive(rm.RoomDriveFunc)
	}
}
