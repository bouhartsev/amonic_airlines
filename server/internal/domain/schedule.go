package domain

import (
	"math"
)

type Schedule struct {
	Id              *int     `json:"id,omitempty"`
	Date            *string  `json:"date,omitempty"`
	Time            *string  `json:"time,omitempty"`
	From            *string  `json:"from,omitempty"`
	To              *string  `json:"to,omitempty"`
	FlightNumber    *int     `json:"flightNumber,omitempty"`
	Aircraft        *string  `json:"aircraft,omitempty"`
	EconomyPrice    *float64 `json:"economyPrice,omitempty"`
	BusinessPrice   *float64 `json:"businessPrice,omitempty"`
	FirstClassPrice *float64 `json:"firstClassPrice,omitempty"`
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
	Date         *string `json:"date,omitempty"`
	Time         *string `json:"time,omitempty"`
	EconomyPrice *string `json:"economyPrice,omitempty"`
}

type ConfirmScheduleRequest struct {
	ScheduleId int `json:"-"`
}

type UnconfirmScheduleRequest struct {
	ScheduleId int `json:"-"`
}
