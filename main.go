package main

import (
	"github.com/gin-gonic/gin"
	"messenger/routes"
)

func main() {
	r := gin.Default()

	routes.InitApi(r)
	routes.InitPages(r)

	r.Run() // listen and serve on 0.0.0.0:8080
}