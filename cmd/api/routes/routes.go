package routes

import (
	"github.com/andersonnerib/ms-golang/cmd/api/controllers"
	"github.com/gin-gonic/gin"
)

func PessoaRoutes(router *gin.Engine) {
	pessoaRoutes := router.Group("/pessoas")

	pessoaRoutes.POST("/", controllers.CreatePessoa)
}
