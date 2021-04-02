package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"zenlink-integral/custom_error"
)

func errorMethodNotExist(ctx *gin.Context){
	errRes := custom_error.InvalidMethod()
	ctx.JSON(http.StatusOK, errRes)
}

func errorParams(msg string, ctx *gin.Context){
	errRes := custom_error.InvalidParams(msg)
	ctx.JSON(http.StatusOK, errRes)
}

func errorDatabaseRetrieve(msg string, ctx *gin.Context){
	errRes := custom_error.DatabaseError(msg)
	ctx.JSON(http.StatusOK, errRes)
}


