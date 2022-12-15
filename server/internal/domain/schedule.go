package domain

import (
	"math"
)

type Schedule struct {
	Id              *int     `json:"id,omitempty" example:"4"`
	Date            *string  `json:"date,omitempty" example:"2017-12-04"`
	Time            *string  `json:"time,omitempty" example:"17:00:00"`
	From            *string  `json:"from,omitempty" example:"AUH"`
	To              *string  `json:"to,omitempty" example:"DOH"`
	FlightNumber    *int     `json:"flightNumber,omitempty" example:"49"`
	Aircraft        *string  `json:"aircraft,omitempty" example:"Boeing 738"`
	EconomyPrice    *float64 `json:"economyPrice,omitempty" example:"12.25"`
	BusinessPrice   *float64 `json:"businessPrice,omitempty" example:"30.20"`
	FirstClassPrice *float64 `json:"firstClassPrice,omitempty" example:"50.0"`
	Confirmed       *bool    `json:"confirmed,omitempty"`
}

func (s *Schedule) CalculatePrices() {
	val := math.Floor(float64(*s.EconomyPrice) * 1.35)
	s.BusinessPrice = &val

	val = math.Floor(float64(*s.BusinessPrice) * 1.3)
	s.FirstClassPrice = &val

	val = math.Floor(*s.EconomyPrice)
	s.EconomyPrice = &val
}

type GetSchedulesRequest struct {
	From         *string `json:"-"`
	To           *string `json:"-"`
	SortBy       *string `json:"-"`
	Outbound     *string `json:"-"`
	FlightNumber *int    `json:"-"`
}

type GetSchedulesResponse struct {
	Schedules []Schedule `json:"schedules"`
}

type UpdateScheduleRequest struct {
	ScheduleId   int     `json:"-"`
	Date         *string `json:"date,omitempty" example:"2017-12-08"`
	Time         *string `json:"time,omitempty" example:"17:00:00"`
	EconomyPrice *string `json:"economyPrice,omitempty" example:"12.25"`
}

type ConfirmScheduleRequest struct {
	ScheduleId int `json:"-"`
}

type UnconfirmScheduleRequest struct {
	ScheduleId int `json:"-"`
}
