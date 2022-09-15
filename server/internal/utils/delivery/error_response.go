package delivery

import (
	"errors"
	"github.com/bouhartsev/amonic_airlines/server/internal/domain/errdomain"
	"github.com/gin-gonic/gin"
	"net/http"
)

func ErrorResponse(c *gin.Context, err error) {
	var (
		e      *errdomain.ErrorResponse
		status int
	)

	if errors.As(err, &e) {
		switch e.Type {
		case errdomain.ObjectNotFoundType:
			status = http.StatusNotFound
		case errdomain.ObjectDuplicateType, errdomain.ObjectDisabledType:
			status = http.StatusConflict
		case errdomain.InvalidRequestType:
			status = http.StatusBadRequest
		case errdomain.InternalType:
			status = http.StatusInternalServerError
		case errdomain.AccessDeniedType:
			status = http.StatusForbidden
		case errdomain.UnauthorizedType:
			status = http.StatusUnauthorized
		default:
			status = http.StatusInternalServerError
		}

		c.JSON(status, e)
		return
	}

	_ = c.AbortWithError(http.StatusInternalServerError, err)
}
