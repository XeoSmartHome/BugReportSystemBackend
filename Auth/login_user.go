package Auth

import (
	"BugReportSystemBackend/Database/Models"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"net/http"
	"time"
)

type signInResponseType struct {
	AccessToken string `json:"accessToken"`
}

type accessTokenClaims struct {
	UserId primitive.ObjectID `json:"userId"`
	jwt.StandardClaims
}

func CreateAccessToken(user Models.User) string {
	now := time.Now()

	claims := accessTokenClaims{
		user.Id,
		jwt.StandardClaims{
			Issuer:    "XeoSmartHome",
			IssuedAt:  now.Unix(),
			NotBefore: now.Unix(),
			ExpiresAt: now.Add(time.Hour * 2).Unix(),
		},
	}
	signingKey := []byte("AllYourBase")
	signedToken, _ := jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString(signingKey)

	return signedToken
}

func SignInUser(ctx *gin.Context, user Models.User) {

	accessToken := CreateAccessToken(user)

	ctx.JSON(http.StatusOK, signInResponseType{
		AccessToken: accessToken,
	})
}
