package server

import (
	"log"
	"net/http"

	"github.com/atbys/koremiyo/interfaces/controller"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

func (s *Server) SessionCheck(ctrl *controller.UserController) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		session := sessions.Default(ctx)
		isLoggedin, userID := ctrl.SessionCheck(session)
		if isLoggedin {
			log.Printf("[+] user_id is %d\n", userID.(int))
			ctx.Set("user_id", userID)
			ctx.Next()
		} else {
			ctx.Redirect(http.StatusFound, "/login")
			ctx.Abort()
		}
	}
}
