package server

import (
	"context"
	"errors"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"go.uber.org/zap"

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
			return []byte(s.cfg.TokenKey), nil
		})

		if _, ok := token.Claims.(*domain.AuthClaims); !ok {
			c.AbortWithStatusJSON(http.StatusUnauthorized, errdomain.InvalidAuthTokenError)
			return
		}

		if err != nil {
			if errors.Is(err, jwt.ErrTokenExpired) {
				err := s.core.SetExpiredTokenError(c.Request.Context(), *claims.User.Id)
				if err != nil {
					s.logger.Error("Failed to set token as expired", zap.Error(err))
				}

				c.AbortWithStatusJSON(http.StatusUnauthorized, errdomain.AuthTokenExpiredError)
				return
			}
			c.AbortWithStatusJSON(http.StatusUnauthorized, errdomain.InvalidAuthTokenError)
			return
		}

		c.Next()
	}
}
