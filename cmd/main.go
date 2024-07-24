package main

import "github.com/gin-gonic/gin"

func main() {
	server := gin.Default()

	v1 := server.Group("/api/v1")
	{
		pets := v1.Group("/pets")
		{
			pets.GET("/")
		}
	}

	server.Run(":8088")
}
