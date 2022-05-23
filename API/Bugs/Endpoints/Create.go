package Endpoints

import (
	"BugReportSystemBackend/Database/Models"
	"BugReportSystemBackend/Database/Repos"
	"fmt"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"net/http"
	"time"
)

func CreateBugsHandler() func(ctx *gin.Context) {

	type Request struct {
		Title       string `json:"title"`
		Description string `json:"description"`
		Status      string `json:"status"`
	}

	type Response struct {
		Bug Models.Bug `json:"bug"`
	}

	return func(ctx *gin.Context) {
		var request Request
		err := ctx.BindJSON(&request)
		fmt.Println(err)
		if err != nil {
			return
		}
		now := primitive.NewDateTimeFromTime(time.Now())
		bug := Models.Bug{
			Title:       request.Title,
			Description: request.Description,
			UserId:      primitive.ObjectID{},
			Status:      request.Status,
			CreatedAt:   now,
			UpdatedAt:   now,
		}
		err = Repos.BugsRepo.InsertOne(&bug)
		if err != nil {
			ctx.AbortWithStatus(http.StatusInternalServerError)
			return
		}
		ctx.JSON(http.StatusOK, Response{
			Bug: bug,
		})
	}
}
