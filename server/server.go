package server

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/savioafs/findAFriendAPI/server/routes"
)

func getStage() string {
	if os.Getenv("STAGE") == "localRun" {
		return ":8088"
	}

	return ":8000"
}

type Server struct {
	port   string
	server *gin.Engine
}

func NewServer() Server {
	return Server{
		port:   getStage(),
		server: gin.Default(),
	}
}

func (s *Server) Run() {
	router := routes.ConfigRoutes(s.server)

	log.Printf("Server running at port: %v", s.port)
	log.Fatal(router.Run(s.port))
}
