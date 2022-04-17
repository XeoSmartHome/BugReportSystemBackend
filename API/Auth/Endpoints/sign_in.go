package Endpoints

import (
	"BugReportSystemBackend/Database/Repos"
	"github.com/gin-gonic/gin"
)

func SignInHandler() func(ctx *gin.Context) {
	type Request struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	return func(ctx *gin.Context) {
		var request Request
		err := ctx.BindJSON(&request)
		if err != nil {
			return
		}
		user, err := Repos.UsersRepo.GetByEmail(request.Email)
		if err != nil {
			return
		}
		if user.PasswordHash != request.Password {
			return
		}

	}
}
