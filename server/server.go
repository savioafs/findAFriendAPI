package server

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"
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
	// define routes here

	log.Printf("Server running at port: %v", s.port)
	// log.Fatal() define fatal route err
}
