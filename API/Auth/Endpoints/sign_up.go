package Endpoints

import (
	"BugReportSystemBackend/API"
	"BugReportSystemBackend/Auth"
	"BugReportSystemBackend/Database/Repos"
	"github.com/gin-gonic/gin"
)

func SignUpHandler() func(ctx *gin.Context) {
	type Request struct {
		FirstName string `json:"firstName"`
		LastName  string `json:"lastName"`
		Email     string `json:"email"`
		Password  string `json:"password"`
	}

	return func(ctx *gin.Context) {
		var request Request
		err := ctx.BindJSON(&request)
		if err != nil {
			return
		}
		user, err := Repos.UsersRepo.CreateUser(request.FirstName, request.LastName, request.Email, request.Password, "PROGRAMMER")
		if err != nil {
			API.SendBadRequest(ctx, "Error")
			return
		}
		Auth.SignInUser(ctx, *user)
	}
}
