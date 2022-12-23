package domain

import (
	"math"
)

type Schedule struct {
	Id              int     `json:"id,omitempty" example:"4"`
	Date            string  `json:"date,omitempty" example:"2017-12-04"`
	Time            string  `json:"time,omitempty" example:"17:00:00"`
	AirportFromId   string  `json:"airportFromId,omitempty" example:"2"`
	AirportToId     string  `json:"airportToId,omitempty" example:"4"`
	FlightNumber    int     `json:"flightNumber,omitempty" example:"49"`
	Aircraft        string  `json:"aircraft,omitempty" example:"Boeing 738"`
	EmptySeats      int     `json:"emptySeats" example:"84"`
	EconomyPrice    float64 `json:"economyPrice,omitempty" example:"12.25"`
	BusinessPrice   float64 `json:"businessPrice,omitempty" example:"30.20"`
	FirstClassPrice float64 `json:"firstClassPrice,omitempty" example:"50.0"`
	Confirmed       bool    `json:"confirmed,omitempty"`
}

func (s *Schedule) CalculatePrices() {
	s.BusinessPrice = math.Floor(float64(s.EconomyPrice) * 1.35)

	s.FirstClassPrice = math.Floor(float64(s.BusinessPrice) * 1.3)

	s.EconomyPrice = math.Floor(s.EconomyPrice)
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
	ScheduleId   int      `json:"-"`
	Date         *string  `json:"date,omitempty" example:"2017-12-08"`
	Time         *string  `json:"time,omitempty" example:"17:00:00"`
	EconomyPrice *float32 `json:"economyPrice,omitempty" example:"12.25"`
}

type SwitchScheduleStatusRequest struct {
	ScheduleId int `json:"-"`
}

type UnconfirmScheduleRequest struct {
	ScheduleId int `json:"-"`
}

type UpdateSchedulesFromFileResponse struct {
	SuccessfulChangesApplied         int `json:"successfulChangesApplied" example:"33"`
	DuplicateRecordsDiscarded        int `json:"duplicateRecordsDiscarded" example:"2"`
	RecordWithMissingFieldsDiscarded int `json:"recordWithMissingFieldsDiscarded" example:"10"`
}

type ScheduleAction struct {
	Action       string
	Date         string
	Time         string
	FlightNumber string
	From         string
	To           string
	Aircraft     string
	Price        string
	Confirmed    string
}
