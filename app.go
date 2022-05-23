package main

import (
	"BugReportSystemBackend/API/Auth"
	"BugReportSystemBackend/API/Bugs"
	"github.com/gin-gonic/gin"
	"net/http"
)

func handleMain(context *gin.Context) {
	context.String(http.StatusOK, "BugReportSystemBackend")
}

func CreateApp() *gin.Engine {
	app := gin.New()
	app.GET("/", handleMain)

	Auth.Init(app)
	Bugs.Init(app)

	return app
}

func main() {
	app := CreateApp()
	_ = app.SetTrustedProxies([]string{"192.168.100.11"})
	_ = app.Run("192.168.100.18:5000")
}
