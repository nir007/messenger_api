package main

import (
	"messenger/routes"
	"os"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.Use(routes.CORS())

	routes.InitSite(r)
	routes.InitREST(r)

	port := os.Getenv("PORT")

	if len(port) == 0 {
		port = ":8081"
	}

	r.Run(port) // listen and serve on 0.0.0.0:8080
}
