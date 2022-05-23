package Bugs

import (
	"BugReportSystemBackend/API/Bugs/Endpoints"
	"github.com/gin-gonic/gin"
)

func Init(app *gin.Engine) {
	bugsRouter := app.Group("/bugs")
	bugsRouter.GET("/bugs", Endpoints.GetBugsHandler())
	bugsRouter.POST("/create-bug", Endpoints.CreateBugsHandler())
}
