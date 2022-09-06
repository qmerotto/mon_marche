package web_api

import (
	"github.com/gin-gonic/gin"
	"test_tech/web_api/handlers"
)

func RunServer() {
	if err := getRouter().Run(":8081"); err != nil {
		panic(err)
	}
}

func getRouter() *gin.Engine {
	r := gin.New()
	r.Use(gin.Recovery())

	v1 := r.Group("/web_api")
	{
		v1.POST("/tickets", func(c *gin.Context) { handlers.Ticket(c) })
	}

	return r
}
