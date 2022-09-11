package utils

import (
	"encoding/json"
	"errors"
	"github.com/bouhartsev/amonic_airlines/server/internal/domain/errdomain"
	"github.com/gin-gonic/gin"
)

func ErrorResponse(c *gin.Context, status int, err error) {
	var errorResponse *errdomain.ErrorResponse

	if errors.As(err, &errorResponse) {
		data, _ := json.Marshal(errorResponse)
		c.Writer.WriteHeader(status)
		_, _ = c.Writer.Write(data)
	} else {
		_ = c.AbortWithError(status, err)
	}
}
