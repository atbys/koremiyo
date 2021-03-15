package server

import (
	"github.com/atbys/koremiyo/infrastructure"
	"github.com/atbys/koremiyo/interfaces/controller"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
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

func (s *Server) Initialize() {
	//gopath := os.Getenv("GOPATH")
	//s.Engine.LoadHTMLGlob(gopath + "/src/github.com/atbys/koremiyo/infrastructure/resource/template/*")
	s.Engine.LoadHTMLGlob("./infrastructure/resource/template/*")
	s.InitializeSession()
	s.SetRouter()
}

var CookieStore cookie.Store

func (s *Server) InitializeSession() {
	CookieStore = cookie.NewStore([]byte("secret"))
	s.Engine.Use(sessions.Sessions("session", CookieStore))
}

func (s *Server) SetRouter() {
	movieController := controller.NewMovieController(infrastructure.NewScraper(), infrastructure.NewCacher())
	userController := controller.NewUserController(infrastructure.NewSqlHandler())

	s.Engine.GET("/", s.showIndex(movieController))
	s.Engine.POST("/random", s.showRandom(movieController))
	s.Engine.GET("/clip/user", s.inputUserFid(movieController))
	s.Engine.GET("/clip/random", s.showRandomFromClip(movieController))
	s.Engine.GET("/user/:id", s.showUser(userController))
	s.Engine.GET("/login", s.showLoginForm(userController))
	s.Engine.POST("/login", s.Login(userController))
	s.Engine.GET("logout", s.Logout(userController))
	s.Engine.GET("/signup", s.showSignUpForm(userController))
	s.Engine.POST("/signup", s.SignUp(userController))

	authGroup := s.Engine.Group("/")
	authGroup.Use(s.SessionCheck(userController))
	{
		authGroup.GET("/loggedin", s.showLoggedin(userController))
		withFriends := authGroup.Group("/with")
		{
			withFriends.GET("/choose", s.showFriends(userController))
			withFriends.POST("/random", s.showMutualClip(movieController))
			withFriends.POST("/majority", s.showMutualClip(movieController))
		}

	}
}

func Run() {
	s := NewServer()
	s.Initialize()
	s.Engine.Run()
}
