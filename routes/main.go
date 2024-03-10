package routes

import (
	"log"
	"strconv"

	"github.com/gin-gonic/gin"
)

var routeBinders = []func(r *gin.Engine){
	BindAuthRoutes,
	BindUserRoutes,
}

type Server struct {
	Port   int
	Engine *gin.Engine
}

func (s *Server) Start() error {
	s.Engine = gin.Default()
	s.Engine.LoadHTMLGlob("./views/*")
	s.BindRoutes()
	log.Println("Server online!")
	err := s.Engine.Run(":" + strconv.Itoa(s.Port))
	return err
}

func (s *Server) BindRoutes() {
	for _, binder := range routeBinders {
		binder(s.Engine)
	}
}
