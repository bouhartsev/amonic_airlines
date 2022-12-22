package server

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	"github.com/bouhartsev/amonic_airlines/server/internal/domain"
	"github.com/bouhartsev/amonic_airlines/server/internal/domain/errdomain"
	"github.com/bouhartsev/amonic_airlines/server/internal/utils/delivery"
)

// CreateUser godoc
// @Summary Создает пользователя.
// @Description Создает пользователя. При дубликате email возвращает ошибку с кодом `user.email:already_taken`.
// @Tags Users
// @Accept json
// @Produce json
// @Param input body domain.CreateUserRequest true "JSON input"
// @Success 201 {object} domain.User
// @Failure 400 {object} errdomain.ErrorResponse
// @Failure 500 {object} errdomain.ErrorResponse
// @Router /api/users [post]
func (s *Server) CreateUser(c *gin.Context) {
	input := new(domain.CreateUserRequest)

	if err := delivery.ReadJson(c.Request, &input); err != nil {
		c.JSON(http.StatusBadRequest, errdomain.InvalidJSONError)
		return
	}

	user, err := s.core.CreateUser(c.Request.Context(), input)

	if err != nil {
		delivery.ErrorResponse(c, err)
		return
	}

	c.JSON(http.StatusCreated, user)
}

// GetUsers godoc
// @Summary Возвращает список пользователей.
// @Tags Users
// @Accept json
// @Produce json
// @Param officeId query int false "Фильтрация по офису, к которому принадлежат пользователи"
// @Success 200 {object} domain.GetUsersResponse
// @Failure 500 {object} errdomain.ErrorResponse
// @Router /api/users [get]
func (s *Server) GetUsers(c *gin.Context) {
	var officeId *int

	if val, err := strconv.Atoi(c.Query("officeId")); err == nil {
		officeId = &val
	}

	users, err := s.core.GetUsers(c.Request.Context(), &domain.GetUsersRequest{OfficeId: officeId})

	if err != nil {
		delivery.ErrorResponse(c, err)
		return
	}

	c.JSON(http.StatusOK, users)
}

// GetUser godoc
// @Summary Возвращает информацию о пользователе.
// @Description Если пользователь не найден, вернет ошибку с кодом `user:not_found`.
// @Tags Users
// @Accept json
// @Produce json
// @Param user_id path int true "Идентификатор пользователя"
// @Success 200 {object} domain.User
// @Failure 404 {object} errdomain.ErrorResponse
// @Failure 500 {object} errdomain.ErrorResponse
// @Router /api/users/{user_id} [get]
func (s *Server) GetUser(c *gin.Context) {
	userId, err := strconv.Atoi(c.Param("user_id"))

	if err != nil {
		delivery.ErrorResponse(c, errdomain.UserNotFoundError)
		return
	}

	user, err := s.core.GetUser(c.Request.Context(), &domain.GetUserRequest{UserId: userId})

	if err != nil {
		delivery.ErrorResponse(c, err)
		return
	}

	c.JSON(http.StatusOK, user)
}

// UpdateUser godoc
// @Summary Возвращает информацию о пользователе.
// @Description Если пользователь не найден, вернет ошибку с кодом `user:not_found`.
// @Tags Users
// @Accept json
// @Produce json
// @Param user_id path int true "Идентификатор пользователя"
// @Param input body domain.UpdateUserRequest true "JSON input"
// @Success 200 {object} domain.User
// @Failure 404 {object} errdomain.ErrorResponse
// @Failure 500 {object} errdomain.ErrorResponse
// @Router /api/users/{user_id} [patch]
func (s *Server) UpdateUser(c *gin.Context) {
	userId, err := strconv.Atoi(c.Param("user_id"))

	if err != nil {
		delivery.ErrorResponse(c, errdomain.UserNotFoundError)
		return
	}

	_, err = s.core.GetUser(c.Request.Context(), &domain.GetUserRequest{UserId: userId})

	if err != nil {
		delivery.ErrorResponse(c, err)
		return
	}

	input := new(domain.UpdateUserRequest)

	if err := delivery.ReadJson(c.Request, &input); err != nil {
		c.JSON(http.StatusBadRequest, errdomain.InvalidJSONError)
		return
	}

	input.UserId = userId
	err = s.core.UpdateUser(c.Request.Context(), input)

	if err != nil {
		c.JSON(http.StatusBadRequest, errdomain.InvalidJSONError)
		return
	}

	updatedUser, err := s.core.GetUser(c.Request.Context(), &domain.GetUserRequest{UserId: userId})

	if err != nil {
		delivery.ErrorResponse(c, err)
		return
	}

	c.JSON(http.StatusOK, updatedUser)
}

// GetUserLogins godoc
// @Summary Возвращает список логинов пользователя.
// @Description Можно запросить список логинов только для одного пользователя.
// @Tags Users
// @Accept json
// @Produce json
// @Param user_id path int true "Id пользователя, для которого запрашиваются логины"
// @Success 200 {object} domain.GetUserLoginsResponse
// @Failure 500 {object} errdomain.ErrorResponse
// @Router /api/users/{user_id}/logins [get]
func (s *Server) GetUserLogins(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("user_id"))

	users, err := s.core.GetUserLogins(c.Request.Context(), id)

	if err != nil {
		delivery.ErrorResponse(c, err)
		return
	}

	c.JSON(http.StatusOK, users)
}

// SwitchUserStatus godoc
// @Summary Переключает флаг active.
// @Tags Users
// @Accept json
// @Produce json
// @Param user_id path int true "Идентификатор расписания"
// @Success 200
// @Failure 500 {object} errdomain.ErrorResponse
// @Router /api/users/{user_id}/switch-status [post]
func (s *Server) SwitchUserStatus(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("user_id"))

	err := s.core.SwitchUserStatus(c.Request.Context(), id)

	if err != nil {
		delivery.ErrorResponse(c, err)
		return
	}

	c.Status(http.StatusOK)
}
