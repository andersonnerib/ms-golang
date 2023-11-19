package main

import (
	controllers "command-line-argumentsC:\\Neri\\Go\\ms-golang\\cmd\\api\\controllers\\create-pessoa.go"
	"github.com/gin-gonic/gin"
)

func PessoaRoutes(router *gin.Engine) {
	pessoaRoutes := router.Group("/pessoas")

	pessoaRoutes.POST("/", controllers.CreatePessoa)
}
