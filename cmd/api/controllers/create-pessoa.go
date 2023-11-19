package controllers

import (
	"net/http"

	use_cases "command-line-argumentsC:\\Neri\\Go\\ms-golang\\internal\\use-cases\\createPessoa.go"
	"github.com/gin-gonic/gin"
)

type createPessoaInput struct {
	name string `json:"name" binding:"required"`
}

func CreatePessoa(ctx *gin.Context) {
	var body createPessoaInput

	if err := ctx.ShouldBindJSON(&body); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"sucess": false,
			"error":  err.Error(),
		})
		return
	}

	useCase := use_cases.NewCreatePessoaUseCase()
	err := useCase.Execute(body.name)
	if err := ctx.ShouldBindJSON(&body); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"sucess": false,
			"error":  err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{
		"sucess": false,
	})
}
