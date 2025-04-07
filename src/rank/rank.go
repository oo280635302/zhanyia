package rank

import (
	"sync"
	"time"
)

var _rankManager RankInfoManager

func init() {
	_rankManager.RankInfoMap = make(map[uint16]*Rank)
}

type RankInfoManager struct {
	RankInfoMap map[uint16]*Rank
	dataLock    sync.RWMutex
}

func (rankMgr *RankInfoManager) GetAllRankType() (typList []uint16) {
	rankMgr.dataLock.RLock()
	defer rankMgr.dataLock.RUnlock()

	for typ := range rankMgr.RankInfoMap {
		typList = append(typList, typ)
	}

	return
}

func (rankMgr *RankInfoManager) Len() int {
	rankMgr.dataLock.RLock()
	defer rankMgr.dataLock.RUnlock()

	return len(rankMgr.RankInfoMap)
}

func (rankMgr *RankInfoManager) HasRankInfo(typ uint16) bool {
	rankMgr.dataLock.RLock()
	defer rankMgr.dataLock.RUnlock()

	_, ok := rankMgr.RankInfoMap[typ]
	return ok
}

func (rankMgr *RankInfoManager) GetRankInfo(typ uint16) *Rank {
	rankMgr.dataLock.Lock()
	defer rankMgr.dataLock.Unlock()

	if rankInfo, ok := rankMgr.RankInfoMap[typ]; ok {
		return rankInfo
	}

	newRankInfo := &Rank{Type: typ}
	newRankInfo.RankList = NewSkipList()

	rankMgr.RankInfoMap[typ] = newRankInfo

	return newRankInfo
}

type Rank struct {
	Type     uint16
	RankList *SkipList

	dataLock sync.RWMutex
}

func (r *Rank) GetRankList(offset, limit uint16) (rankList []RankElemInfo) {
	r.dataLock.RLock()
	defer r.dataLock.RUnlock()

	start, end := offset+1, offset+limit
	return r.RankList.Range(int32(start), int32(end))
}

func (r *Rank) SetScore(id, lid, score int64, roleHeader StRoleHeader, leagueHeader StLeagueHeader, modifyTime time.Time) {
	if score <= 0 {
		return
	}
	r.dataLock.Lock()
	defer r.dataLock.Unlock()

	elem := NewSelfRankElem(r.Type, id, lid, score, roleHeader, leagueHeader, modifyTime)
	r.RankList.Set(id, elem, score)
}

func (r *Rank) GetRankElem(id int64) (rankElem RankElemInfo, ok bool) {
	r.dataLock.RLock()
	defer r.dataLock.RUnlock()

	rankElem, ok = r.RankList.Get(id)
	return
}

func (r *Rank) RemoveRankInfo(id int64) {
	r.dataLock.Lock()
	defer r.dataLock.Unlock()

	r.RankList.Remove(id)
}

func (r *Rank) ClearRank() {
	r.dataLock.Lock()
	defer r.dataLock.Unlock()

	r.RankList = NewSkipList()
}

// ///
type RankElemInfo struct {
	Type         uint16
	Id           int64
	Lid          int64
	RoleHeader   StRoleHeader   `ignore:"true"`
	LeagueHeader StLeagueHeader `ignore:"true"`
	Score        int64
	CacheRank    int32 `ignore:"true"`
	CreateAt     time.Time
}

type StRoleHeader struct {
	Icon uint32
	Name string
}

type StLeagueHeader struct {
	Name      string
	ShortName string
	Flag      string
}

func NewSelfRankElem(typ uint16, id, lid, score int64, roleHeader StRoleHeader, leagueHeader StLeagueHeader, modifyTime time.Time) RankElemInfo {
	if score < 0 {
		score = 0
	}

	return RankElemInfo{
		Type:         typ,
		Id:           id,
		Lid:          lid,
		RoleHeader:   roleHeader,
		LeagueHeader: leagueHeader,
		Score:        score,
		CreateAt:     modifyTime,
	}
}
