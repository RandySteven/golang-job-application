package middleware

import (
	"job-application/apperror"
	"job-application/configs"
	"job-application/entity/payload"
	"job-application/utils"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

func validateToken(c *gin.Context) *configs.JWTClaim {
	tokenString, err := c.Cookie("token")
	log.Println("token str : ", tokenString)
	if err != nil {
		return nil
	}
	// tokenStringValidate := c.GetHeader("Authorization")

	// if tokenStringValidate == "" {
	// 	return nil
	// }

	claims := &configs.JWTClaim{}
	token, err := jwt.ParseWithClaims(tokenString, claims, func(t *jwt.Token) (interface{}, error) {
		return configs.JWT_KEY, nil
	})

	if err != nil || !token.Valid {
		return nil
	}

	return claims
}

func AuthMiddleware(c *gin.Context) {
	if os.Getenv("ENV_MODE") == "testing" {
		c.Next()
		return
	}
	claims := validateToken(c)

	if claims == nil {
		resp := payload.Response{
			Errors: []string{apperror.NewErrUnauthorized().Error()},
		}
		utils.ResponseHandler(c.Writer, http.StatusUnauthorized, resp)
		c.Abort()
		return
	}

	c.Set("id", claims.ID)
	c.Set("name", claims.Name)
	c.Set("email", claims.Email)
	c.Next()
}
