package middleware

import (
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"student-halls.com/internal/api"
	"student-halls.com/internal/config"
)

var jwtSecret = []byte(config.AppConfig().Auth.JWTSecret)

func GenerateJWTToken(email string, username string) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)

	// Create a map to store our claims
	claims := token.Claims.(jwt.MapClaims)

	// Set token claims
	claims["iss"] = config.AppConfig().Auth.JWTIssuer
	claims["email"] = email
	claims["username"] = username
	claims["exp"] = time.Now().Add(time.Hour * time.Duration(config.AppConfig().Auth.JWTExpireInHours)).Unix()
	claims["iat"] = time.Now().Unix()

	tokenString, err := token.SignedString(jwtSecret)
	return tokenString, err
}

// JWTAuthMiddleware is a middleware to authenticate the user using JWT
func JWTAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		const BearerSchema = "Bearer "
		authHeader := c.GetHeader("Authorization")

		if !strings.HasPrefix(authHeader, BearerSchema) {
			api.Error(c, http.StatusUnauthorized, "Authorization header must start with Bearer")
			return
		}

		tokenString := authHeader[len(BearerSchema):]
		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
			}
			return []byte(jwtSecret), nil
		})

		if err != nil {
			var errMsg string
			if ve, ok := err.(*jwt.ValidationError); ok {
				if ve.Errors&jwt.ValidationErrorMalformed != 0 {
					errMsg = "Malformed token"
				} else if ve.Errors&jwt.ValidationErrorExpired != 0 {
					errMsg = "Token is expired"
				} else if ve.Errors&jwt.ValidationErrorNotValidYet != 0 {
					errMsg = "Token not valid yet"
				} else {
					errMsg = "Invalid token"
				}
			} else {
				errMsg = "Invalid token"
			}
			api.Error(c, http.StatusUnauthorized, errMsg)
			return
		}

		if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
			c.Set("email", claims["email"])
			c.Set("accountType", claims["accountType"])
		} else {
			api.Error(c, http.StatusUnauthorized, "Invalid token")
			return
		}

		c.Next()
	}
}
