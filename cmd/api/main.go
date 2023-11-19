package main

import (
	routes "github.com/andersonnerib/ms-golang/cmd/api/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.GET("/healthy", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"sucess": true,
		})
	})

	routes.PessoaRoutes(router)

	router.Run(":8080") // listen and serve on 0.0.0.0:8080
}
