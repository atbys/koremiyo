package infrastructure

import (
	"os"

	"github.com/atbys/koremiyo/interfaces/controller"
	"github.com/gin-gonic/gin"
)

type Server struct {
	Engine  *gin.Engine
	confing string //TODO
}

func NewServer() *Server {
	s := &Server{
		Engine: gin.Default(),
	}

	return s
}

func (s *Server) showIndex(ctrl *controller.MovieController) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		res_code, res_data := ctrl.Index()
		ctx.HTML(res_code, "index.html", gin.H{ //res_dataを直接突っ込めないか
			//レンダリング用の関数を作って，KeyとValueを取得するループで突っ込んでいこう
			"title":     res_data.Content["page_title"],
			"recommend": res_data.Content["movie_title"],
		})
	}
}

func (s *Server) showRandom(ctrl *controller.MovieController) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		res_code, res_data := ctrl.Random()
		ctx.HTML(res_code, "movie.html", gin.H{
			"title":       "kore",
			"movie_title": res_data.Content["movie_title"],
		})
	}
}

func (s *Server) showRandomFromClip(ctrl *controller.MovieController) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		res_code, res_data := ctrl.RandomClip("nekoneon") //XXX
		ctx.HTML(res_code, "movie.html", gin.H{
			"title":       "kore",
			"movie_title": res_data.Content["movie_title"],
		})
	}
}

func (s *Server) SetRouter() {
	movieController := controller.NewMovieController(NewScraper())
	gopath := os.Getenv("GOPATH")
	s.Engine.LoadHTMLGlob(gopath + "/src/github.com/atbys/koremiyo/infrastructure/resource/template/*")

	s.Engine.GET("/", s.showIndex(movieController))
	s.Engine.GET("/random", s.showRandom(movieController))
	s.Engine.GET("/clip", s.showRandomFromClip(movieController))
}

func Run() {
	s := NewServer()
	s.SetRouter()
	s.Engine.Run()
}
