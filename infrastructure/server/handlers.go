package server

import (
	"log"
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

func (s *Server) showMutualClip(ctrl *controller.MovieController) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ids := ctx.PostFormArray("filmarks_id")
		for _, id := range ids {
			println(id)
		}
		code, data := ctrl.MutualClip(ids)
		ctx.HTML(code, "movie.html", gin.H{
			"title":       "kore",
			"movie_title": data.Movie.Title,
			"movie_rate":  data.Content["movie_rate"],
			"reviews":     data.Movie.Reviews,
			"link":        data.Content["link"],
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

func (s *Server) showFriends(ctrl *controller.UserController) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		v, _ := ctx.Get("user_id")
		userID := v.(int)
		log.Printf("[+] Got user_id is %d\n", userID)
		code, users := ctrl.SelectFriend(userID)

		ctx.HTML(code, "select_friends.html", gin.H{
			"friends": users,
		})
	}
}
