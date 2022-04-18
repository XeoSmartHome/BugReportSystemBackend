package API

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type BadRequest struct {
	Error string `json:"error"`
}

func SendBadRequest(ctx *gin.Context, error string) {
	ctx.JSON(http.StatusBadRequest, BadRequest{Error: error})
}
