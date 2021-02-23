package usecase

import (
	"math/rand"
	"sort"
	"time"

	"github.com/atbys/koremiyo/domain"
)

type MovieInteractor struct {
	MovieRepository  MovieRepository
	MovieOutputPort  MovieOutputPort
	MutualMovieCache MutualMovieCache
}

type OutputData struct {
	CacheID int
	Movie   *domain.Movie
	Msg     map[string]interface{}
}

type OutputConfig struct{}

const allMovieNum = 85000

func (interactor *MovieInteractor) GetRecommendation() (*OutputData, error) { //TODO
	movie := &domain.Movie{
		Title: "1917",
	}
	return interactor.MovieOutputPort.ShowIndex(movie)
}

// GetMovieInfo : 指定されたIDの映画を取得してくる
func (interactor *MovieInteractor) GetMovieInfo(id int) (*OutputData, error) {
	movie, err := interactor.MovieRepository.FindById(id)
	if err != nil {
		return nil, err
	}
	return interactor.MovieOutputPort.ShowMovieInfo(movie)
}

func randInt(max int) int {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(max)
}

//RandomNumberGenerator 指定された範囲内のランダムな数値を出力する
type RandomNumberGenerator interface {
	Intn(int) int
}

// RNG (RandomNumberGenerator) 指定された範囲内のランダムな数値を出力する
var RNG RandomNumberGenerator = rand.New(rand.NewSource(time.Now().UnixNano()))

// GetRandom : Filmarksに記録されている映画から適当に選出する
func (interactor *MovieInteractor) GetRandom() (*OutputData, error) {
	randID := RNG.Intn(allMovieNum)
	movie, err := interactor.MovieRepository.FindById(randID)
	if err != nil {
		return nil, err
	}
	return interactor.MovieOutputPort.ShowMovieInfo(movie)
}

// GetRandomFromClips : FilmarksユーザIDからClipした映画をとってきてランダムに選出
func (interactor *MovieInteractor) GetRandomFromClips(userID string) (*OutputData, error) {
	movieIDs, _ := interactor.MovieRepository.FindClipsByUserId(userID)
	randID := movieIDs[RNG.Intn(len(movieIDs))]
	movie, _ := interactor.MovieRepository.FindById(randID)
	return interactor.MovieOutputPort.ShowMovieInfo(movie)
}

func (interactor *MovieInteractor) GetMutualClip(filmarksIDs []string, cacheID int) (*OutputData, error) {
	var allMutualClip map[int]int
	if cacheID < 0 {
		allMutualClip = make(map[int]int)
		for _, fid := range filmarksIDs {
			clipMovies, _ := interactor.MovieRepository.FindClipsByUserId(fid)
			for _, mid := range clipMovies {
				setDefault(allMutualClip, mid, 0)
				allMutualClip[mid]++
			}
		}
		mutualRank := List{}
		for k, v := range allMutualClip {
			e := Entry{k, v}
			mutualRank = append(mutualRank, e)
		}
		sort.Sort(mutualRank)
		cacheID := interactor.MutualMovieCache.Store(mutualRank, 1)
		movie, _ := interactor.MovieRepository.FindById(mutualRank[0].mid)
		out, err := interactor.MovieOutputPort.ShowMovieInfo(movie)
		out.CacheID = cacheID
		out.Msg["cache_id"] = cacheID
		return out, err
	} else {
		mutualRank, index, err := interactor.MutualMovieCache.FindById(cacheID)
		if err != nil {
			return nil, err
		}
		movie, err := interactor.MovieRepository.FindById(mutualRank[index].mid)
		cacheID := interactor.MutualMovieCache.Store(mutualRank, index+1)
		out, err := interactor.MovieOutputPort.ShowMovieInfo(movie)
		out.CacheID = cacheID
		out.Msg["cache_id"] = cacheID
		if err != nil {
			return nil, err
		}
		return out, nil
	}
}

func setDefault(m map[int]int, key, value int) {
	if _, ok := m[key]; !ok {
		m[key] = value
	}
}

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
