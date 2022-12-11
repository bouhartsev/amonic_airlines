package server

import (
	"context"
	"fmt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"

	"github.com/bouhartsev/amonic_airlines/server/internal/domain"
	"github.com/bouhartsev/amonic_airlines/server/internal/domain/errdomain"
)

func getTokenMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		header := c.Request.Header.Get("Authorization")
		c.Request = c.Request.WithContext(context.WithValue(c.Request.Context(), "token", header))
		c.Next()
	}
}

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PATCH, PUT, DELETE")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}

func (s *Server) checkAuthorizationMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		p := c.Request.URL.Path

		if p == `/api/auth/sign-in` || strings.HasPrefix(p, `/api/docs`) {
			c.Next()
			return
		}

		header := c.Request.Header.Get("Authorization")

		if header == "" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, errdomain.AuthTokenIsNotPresentedError)
			return
		}

		vals := strings.Split(header, "Bearer ")

		if len(vals) != 2 {
			c.AbortWithStatusJSON(http.StatusUnauthorized, errdomain.InvalidAuthTokenError)
			return
		}

		accessToken := vals[1]

		claims := &domain.AuthClaims{}

		token, err := jwt.ParseWithClaims(accessToken, claims, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
			}

			return []byte(s.cfg.TokenKey), nil
		})

		if err != nil {
			fmt.Println("here 1", err.Error())
			c.AbortWithStatusJSON(http.StatusUnauthorized, errdomain.InvalidAuthTokenError)
			return
		}

		if _, ok := token.Claims.(*domain.AuthClaims); !ok && !token.Valid {
			fmt.Println("here 4")
			c.AbortWithStatusJSON(http.StatusUnauthorized, errdomain.InvalidAuthTokenError)
			return
		}

		c.Next()
	}
}
