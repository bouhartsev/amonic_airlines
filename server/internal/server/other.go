package server

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/bouhartsev/amonic_airlines/server/internal/utils/delivery"
)

// GetAirports godoc
// @Summary Возвращает список аэропортов.
// @Tags Airports
// @Accept json
// @Produce json
// @Success 200 {object} domain.GetAirportsResponse
// @Failure 500 {object} errdomain.ErrorResponse
// @Router /api/airports [get]
func (s *Server) GetAirports(c *gin.Context) {
	airports, err := s.core.GetAirports(c.Request.Context())

	if err != nil {
		delivery.ErrorResponse(c, err)
		return
	}

	c.JSON(http.StatusOK, airports)
}

// GetCabinTypes godoc
// @Summary Возвращает список типов cabins.
// @Tags CabinTypes
// @Accept json
// @Produce json
// @Success 200 {object} domain.GetCabinTypesResponse
// @Failure 500 {object} errdomain.ErrorResponse
// @Router /api/cabin-types [get]
func (s *Server) GetCabinTypes(c *gin.Context) {
	cb, err := s.core.GetCabinTypes(c.Request.Context())

	if err != nil {
		delivery.ErrorResponse(c, err)
		return
	}

	c.JSON(http.StatusOK, cb)
}

// GetCountries godoc
// @Summary Возвращает список стран.
// @Tags Countries
// @Accept json
// @Produce json
// @Success 200 {object} domain.GetCountriesResponse
// @Failure 500 {object} errdomain.ErrorResponse
// @Router /api/countries [get]
func (s *Server) GetCountries(c *gin.Context) {
	cb, err := s.core.GetCountries(c.Request.Context())

	if err != nil {
		delivery.ErrorResponse(c, err)
		return
	}

	c.JSON(http.StatusOK, cb)
}

// GetOffices godoc
// @Summary Возвращает список офисов.
// @Tags Offices
// @Accept json
// @Produce json
// @Success 200 {object} domain.GetOfficesResponse
// @Failure 500 {object} errdomain.ErrorResponse
// @Router /api/offices [get]
func (s *Server) GetOffices(c *gin.Context) {
	response, err := s.core.GetOffices(c.Request.Context())

	if err != nil {
		delivery.ErrorResponse(c, err)
		return
	}

	c.JSON(http.StatusOK, response)
}
