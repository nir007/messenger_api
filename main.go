package main

import (
	"messenger/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	routes.InitAPI(r)
	routes.InitPages(r)

	r.Run(":8082") // listen and serve on 0.0.0.0:8080
}
