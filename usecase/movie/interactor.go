package movie

import (
	"math/rand"
	"sort"
	"time"

	"github.com/atbys/koremiyo/domain"
)

type MovieInteractor struct {
	Repository MovieRepository
	Cache      MovieCache
}

const allMovieNum = 85000

//RandomNumberGenerator 指定された範囲内のランダムな数値を出力する
type RandomNumberGenerator interface {
	Intn(int) int
}

// RNG (RandomNumberGenerator) 指定された範囲内のランダムな数値を出力する
var RNG RandomNumberGenerator = rand.New(rand.NewSource(time.Now().UnixNano()))

func (mi *MovieInteractor) Random() (*domain.Movie, error) {
	mid := RNG.Intn(allMovieNum)
	m, err := mi.Repository.Find(mid)
	if err != nil {
		return nil, err
	}

	return m, nil
}

func (mi *MovieInteractor) RandomFromClips(fid string) (*domain.Movie, error) {
	mids, err := mi.Repository.FindClips(fid)
	if err != nil {
		return nil, err
	}

	randMID := mids[RNG.Intn(len(mids))]
	m, _ := mi.Repository.Find(randMID)

	return m, nil
}

func (mi *MovieInteractor) RandomFromMutual(fids []string, cacheID int) (*domain.Movie, int, error) {
	if cacheID < 0 {
		mc := NewMutualClips(fids)
		mc.Init(mi)

		cacheID = mi.Cache.Store(mc)

		mid := mc.Random()
		m, err := mi.Repository.Find(mid)
		if err != nil {
			return nil, cacheID, err
		}

		return m, cacheID, nil
	} else {
		cachedData, err := mi.Cache.Find(cacheID)
		if err != nil {
			return nil, cacheID, err
		}
		mc := cachedData.(*MutualClips)

		mid := mc.Random()
		m, err := mi.Repository.Find(mid)
		if err != nil {
			return nil, 0, err
		}

		return m, cacheID, nil
	}
}

func (mi *MovieInteractor) MajorityFromMutual(fids []string, cacheID int) (*domain.Movie, int, error) {
	if cacheID < 0 {
		mc := NewMutualClips(fids)
		mc.Init(mi)

		mid := mc.Majority()

		cacheID = mi.Cache.Store(mc)

		m, err := mi.Repository.Find(mid)
		if err != nil {
			return nil, cacheID, err
		}

		return m, cacheID, nil
	} else {
		cachedData, err := mi.Cache.Find(cacheID)
		if err != nil {
			return nil, cacheID, err
		}
		mc := cachedData.(*MutualClips)

		mid := mc.Majority()
		m, err := mi.Repository.Find(mid)
		if err != nil {
			return nil, cacheID, err
		}

		return m, cacheID, nil
	}
}

//=====MutuakClipsについて==========
type MutualClips struct {
	Participants []string
	Clips        map[int][]string
	Marks        map[int][]string
	Ranked       []int
	index        int
	CacheID      int
}

func NewMutualClips(member []string) *MutualClips {
	mc := &MutualClips{
		Participants: member,
		Clips:        make(map[int][]string),
		Ranked:       nil,
		index:        0,
	}

	return mc
}

func (mc *MutualClips) Init(mi *MovieInteractor) {
	for _, fid := range mc.Participants {
		clipMovies, _ := mi.Repository.FindClips(fid)
		for _, mid := range clipMovies {
			mc.Clips[mid] = append(mc.Clips[mid], fid) // だれがClipしているのかを記録
		}
	}
	//TODO: Markしている映画も取得して記録する
}

func (mc *MutualClips) Random() (mid int) { //TODO: だれがClipした映画なのかも返す
	rn := RNG.Intn(len(mc.Clips))
	index := 0
	for key, _ := range mc.Clips {
		if index == rn {
			return key
		} else {
			index++
		}
	}

	return -1
}

func (mc *MutualClips) SortRank() {
	mc.Ranked = make([]int, 0)
	mutualRank := List{}
	for k, v := range mc.Clips {
		counts := len(v)
		e := Entry{k, counts}
		mutualRank = append(mutualRank, e)
	}
	sort.Sort(mutualRank)
	for _, e := range mutualRank {
		mc.Ranked = append(mc.Ranked, e.mid)
	}
}

func (mc *MutualClips) Majority() (mid int) {
	if mc.Ranked == nil {
		mc.SortRank()
	}
	mid = mc.Ranked[mc.index]
	mc.index++

	return
}

//======ソートに使うためのもの=========
type Entry struct {
	mid    int
	counts int
}
type List []Entry

func (l List) Len() int {
	return len(l)
}

func (l List) Swap(i, j int) {
	l[i], l[j] = l[j], l[i]
}

func (l List) Less(i, j int) bool {
	if l[i].counts == l[j].counts {
		return (l[i].mid > l[j].mid)
	}
	return (l[i].counts > l[j].counts)
}
