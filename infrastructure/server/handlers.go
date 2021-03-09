package server

import (
	"log"
	"net/http"
	"strconv"

	"github.com/atbys/koremiyo/domain"
	"github.com/atbys/koremiyo/interfaces/controller"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

func (s *Server) showIndex(ctrl *controller.MovieController) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		code, data := ctrl.Index()
		ctx.HTML(code, "index.html", msgWriter(data.Msg))
	}
}

func (s *Server) showRandom(ctrl *controller.MovieController) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		code, data := ctrl.Random()
		ctx.HTML(code, "movie.html", msgWriter(data.Msg))
	}
}

func (s *Server) inputUserFid(ctrl *controller.MovieController) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.HTML(http.StatusOK, "input_user.html", gin.H{})
	}
}

func (s *Server) showRandomFromClip(ctrl *controller.MovieController) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		fid := ctx.Query("fid")
		println(fid)
		code, data := ctrl.RandomClip(fid)
		ctx.HTML(code, "movie.html", msgWriter(data.Msg))
	}
}

func (s *Server) showMutualClip(ctrl *controller.MovieController) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ids := ctx.PostFormArray("filmarks_id")
		rawCacheID := ctx.PostForm("cache")
		cache, _ := strconv.Atoi(rawCacheID)
		println("cahceID: ", rawCacheID)
		for _, id := range ids {
			println(id)
		}
		code, data := ctrl.MutualClip(ids, cache)
		ctx.HTML(code, "movie.html", msgWriter(data.Msg))
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

func (s *Server) SignUp(ctrl *controller.UserController) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		fid := ctx.PostForm("fid")
		scrName := ctx.PostForm("screen")
		password := ctx.PostForm("pass")
		u := domain.User{
			FilmarksID: fid,
			ScreenName: scrName,
			Password:   password,
		}

		ctrl.SignUp(u)
	}
}

func (s *Server) showSignUpForm(ctrl *controller.UserController) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.HTML(http.StatusOK, "signup.html", gin.H{})
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

func msgWriter(msg map[string]interface{}) (h gin.H) {
	h = gin.H{}

	for k, v := range msg {
		if s, ok := v.(string); ok {
			h[k] = s
		} else if ss, ok := v.([]string); ok {
			h[k] = ss
		} else {
			h[k] = v
		}
	}
	return
}
