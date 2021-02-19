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

type MovieOutputPort interface {
	ShowMovieInfo(*domain.Movie) (*OutputData, error)
	ShowIndex(*domain.Movie) (*OutputData, error)
}

type OutputData struct {
	CacheID int
	Config  *OutputConfig
	Movie   *domain.Movie
	Content map[string]string
}

type OutputConfig struct{}

const allMovieNum = 85000

func (interactor *MovieInteractor) GetRecommendation() (*OutputData, error) {
	movie := &domain.Movie{
		Title: "1917",
	}
	return interactor.MovieOutputPort.ShowIndex(movie)
}

func (interactor *MovieInteractor) GetMovieInfo(id int) (*OutputData, error) {
	movie, err := interactor.MovieRepository.FindById(id)
	if err != nil {
		panic(err)
	}
	return interactor.MovieOutputPort.ShowMovieInfo(movie)
}

func randInt(max int) int {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(max)
}

func (interactor *MovieInteractor) GetRandom() (*OutputData, error) {
	randId := randInt(allMovieNum)
	movie, _ := interactor.MovieRepository.FindById(randId)
	return interactor.MovieOutputPort.ShowMovieInfo(movie)
}

func (interactor *MovieInteractor) GetRandomFromClips(userId string) (*OutputData, error) {
	movieIDs, _ := interactor.MovieRepository.FindByUserId(userId)
	randId := movieIDs[randInt(len(movieIDs))]
	movie, _ := interactor.MovieRepository.FindById(randId)
	return interactor.MovieOutputPort.ShowMovieInfo(movie)
}

func (interactor *MovieInteractor) GetMutualClip(filmarksIDs []string, cacheID int) (*OutputData, error) {
	var allClipMovies map[int]int
	if cacheID < 0 {
		allClipMovies = make(map[int]int)
		for _, fid := range filmarksIDs {
			clipMovies, _ := interactor.MovieRepository.FindByUserId(fid)
			for _, mid := range clipMovies {
				setDefault(allClipMovies, mid, 0)
				allClipMovies[mid] += 1
			}
		}

		a := List{}
		for k, v := range allClipMovies {
			e := Entry{k, v}
			a = append(a, e)
		}
		sort.Sort(a)
		cacheID = interactor.MutualMovieCache.Store(a, 0)
		movie, _ := interactor.MovieRepository.FindById(a[0].mid)
		out, err := interactor.MovieOutputPort.ShowMovieInfo(movie)
		out.CacheID = cacheID
		return out, err
	}

	a, index, _ := interactor.MutualMovieCache.FindById(cacheID)

	movie, _ := interactor.MovieRepository.FindById(a[index].mid)
	cacheID = interactor.MutualMovieCache.Store(a, index+1) //Updateの実装をする
	out, err := interactor.MovieOutputPort.ShowMovieInfo(movie)
	out.CacheID = cacheID
	return out, err
}

func setDefault(m map[int]int, key, value int) {
	if _, ok := m[key]; !ok {
		m[key] = value
	}
}

type Entry struct {
	mid   int
	value int
}
type List []Entry

func (l List) Len() int {
	return len(l)
}

func (l List) Swap(i, j int) {
	l[i], l[j] = l[j], l[i]
}

func (l List) Less(i, j int) bool {
	return (l[i].value > l[j].value)
}
