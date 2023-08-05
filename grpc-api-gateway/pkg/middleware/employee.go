package middleware

import (
	"fmt"
	"net/http"
	"time"

	"github.com/ashiqsabith123/shiftsync-grpc-api-gateway/pkg/utils"
	"github.com/gin-gonic/gin"
)

func CheckSignupCookie(ctx *gin.Context) {
	cookie, err := ctx.Cookie("_signup-cookie")

	if err == nil || cookie != "" {
		ctx.AbortWithStatusJSON(http.StatusAlreadyReported, gin.H{
			"status code": "208",
			"message":     "otp is already send to you mobile verify otp or try after some time",
		})
		return
	}

	ctx.Next()
}

func CheckCookieIsPresent(ctx *gin.Context) {
	cookie, err := ctx.Cookie("_signup-cookie")

	if err != nil || cookie == "" {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"status code": "400",
			"message":     "No details found",
		})
		return
	}

	ctx.Next()
}

func AuthenticateUser(ctx *gin.Context) {
	authtoken(ctx, "_employee-cookie")
}

func authtoken(ctxt *gin.Context, user string) {
	token, err := ctxt.Cookie(user)

	if err != nil || token == "" {
		ctxt.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"status code ": 401,
			"msg":          "Unauthorized User Please Login No token found",
			"err":          fmt.Sprint(err),
		})
		return
	}

	claims, err := utils.ValidateTokens(token)

	if err != nil {
		ctxt.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"StatusCode": 401,
			"msg":        "Unauthorized User Please Login",
			"err":        fmt.Sprint(err),
		})
		return
	}

	if time.Now().Unix() > claims.ExpiresAt {
		ctxt.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"StatusCode": 401,
			"msg":        "User Need Re-Login time expired",
		})
		return
	}

	// claim the userId and set it on context
	ctxt.Set("userId", claims.Id)
}
