package infrastructure

import (
	"net/http"
	"strconv"

	"github.com/atbys/koremiyo/interfaces/controller"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

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
			"movie_rate":  res_data.Content["movie_rate"],
			"reviews":     res_data.Movie.Reviews,
			"link":        res_data.Content["link"],
		})
	}
}

func (s *Server) showUser(ctrl *controller.UserController) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id, _ := strconv.Atoi(ctx.Param("id"))
		code, data := ctrl.Show(id)
		ctx.HTML(code, "user_info.html", gin.H{
			"title":     "user",
			"user_name": data.User.ScreenName,
			"user_fid":  data.User.FilmarksID,
		})
	}
}

func (s *Server) showLoginForm(ctrl *controller.UserController) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.HTML(http.StatusOK, "login.html", gin.H{})
	}
}
func (s *Server) Login(ctrl *controller.UserController) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		session := sessions.Default(ctx)
		userID := ctx.PostForm("user_id")
		password := ctx.PostForm("password")
		code := ctrl.Login(userID, password, session)
		if code == http.StatusBadGateway { //TODO
			ctx.HTML(code, "error.html", gin.H{})
		}
		ctx.Redirect(http.StatusFound, "/loggedin")
	}
}

func (s *Server) Logout(ctrl *controller.UserController) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		session := sessions.Default(ctx)
		code := ctrl.Logout(session)
		ctx.HTML(code, "index.html", gin.H{})
	}
}

func (s *Server) showLoggedin(ctrl *controller.UserController) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.HTML(http.StatusOK, "loggedin.html", gin.H{})
	}
}
