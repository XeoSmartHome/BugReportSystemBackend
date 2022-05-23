package Endpoints

import (
	"BugReportSystemBackend/Database/Models"
	"BugReportSystemBackend/Database/Repos"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetBugsHandler() func(ctx *gin.Context) {

	type Response struct {
		Bugs []Models.Bug `json:"bugs"`
	}

	return func(ctx *gin.Context) {
		bugs, err := Repos.BugsRepo.GetAll()
		if err != nil {
			fmt.Println(err)
			ctx.AbortWithStatus(http.StatusInternalServerError)
			return
		}
		ctx.JSON(http.StatusOK, Response{
			Bugs: bugs,
		})
	}
}
