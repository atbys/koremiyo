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
		res := ctrl.Index()
		ctx.HTML(res.Status, "index.html", msgWriter(res.Message))
	}
}

func (s *Server) showRandom(ctrl *controller.MovieController) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		res := ctrl.Random()
		ctx.HTML(res.Status, "movie.html", msgWriter(res.Message))
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
		res := ctrl.RandomClip(fid)
		ctx.HTML(res.Status, "movie.html", msgWriter(res.Message))
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
		res := ctrl.RandomMutual(ids, cache)
		ctx.HTML(res.Status, "movie.html", msgWriter(res.Message))
	}
}

func (s *Server) showUser(ctrl *controller.UserController) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id := ctx.Param("id")
		res := ctrl.ShowInfo(id)
		ctx.HTML(res.Status, "user_info.html", msgWriter(res.Message))
	}
}

func (s *Server) SignUp(ctrl *controller.UserController) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		fid := ctx.PostForm("fid")
		scrName := ctx.PostForm("screen")
		password := ctx.PostForm("pass")

		res := ctrl.Create(scrName, fid, password)
		ctx.HTML(res.Status, "user_info.html", msgWriter(res.Message))
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
		res := ctrl.Login(userID, password, session)
		if res.Status != http.StatusOK {
			ctx.HTML(res.Status, "error.html", gin.H{})
		}
		ctx.Redirect(http.StatusFound, "/loggedin")
	}
}

func (s *Server) Logout(ctrl *controller.UserController) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		session := sessions.Default(ctx)
		res := ctrl.Logout(session)
		ctx.HTML(res.Status, "index.html", gin.H{}) // TODO: ログアウトページに遷移させる？
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
		res := ctrl.ListFriends(userID)

		ctx.HTML(res.Status, "select_friends.html", msgWriter(res.Message)) //TODO: htmlをいい感じに書き換える
	}
}

func (s *Server) inputFriend(ctrl *controller.UserController) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.HTML(http.StatusOK, "follow_friend.html", gin.H{})
	}
}

func (s *Server) FollowFriend(ctrl *controller.UserController) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		v, _ := ctx.Get("user_id")
		uid := v.(int)
		friendID := ctx.PostForm("friend_id")
		res := ctrl.FollowFriend(uid, friendID)
		if res.Status != http.StatusOK {
			ctx.HTML(res.Status, "error.html", msgWriter(res.Message))
			return
		}
		ctx.HTML(res.Status, "index.html", gin.H{})

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
