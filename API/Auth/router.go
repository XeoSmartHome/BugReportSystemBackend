package Auth

import (
	"BugReportSystemBackend/API/Auth/Endpoints"
	"github.com/gin-gonic/gin"
)

func Init(app *gin.Engine) {
	authRouter := app.Group("/auth")
	authRouter.POST("/sign-in", Endpoints.SignInHandler())
	authRouter.POST("/sign-up", Endpoints.SignUpHandler())
}
