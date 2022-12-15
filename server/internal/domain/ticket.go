package domain

type Ticket struct {
	Id                *int    `json:"id,omitempty"`
	UserId            *int    `json:"userId,omitempty"`
	ScheduleId        *int    `json:"scheduleId,omitempty"`
	CabinTypeId       *int    `json:"cabinTypeId,omitempty"`
	FirstName         *string `json:"firstName,omitempty"`
	LastName          *string `json:"lastName,omitempty"`
	Phone             *string `json:"phone,omitempty"`
	PassportNumber    *string `json:"passportNumber,omitempty"`
	PassportCountryId *int    `json:"passportCountryId,omitempty"`
	BookingReference  *string `json:"bookingReference,omitempty"`
	Confirmed         *bool   `json:"confirmed,omitempty"`
}

type AddTicketRequest struct {
	Outbound  TicketInfo  `json:"outbound"`
	Return    *TicketInfo `json:"return,omitempty"`
	Passenger struct {
		Firstname         string `json:"firstname"`
		Lastname          string `json:"lastname"`
		Birthdate         string `json:"birthdate"`
		PassportNumber    string `json:"passportNumber"`
		PassportCountryId int    `json:"passportCountryId"`
		Phone             string `json:"phone"`
	}
}

type TicketInfo struct {
	ScheduleId  int `json:"scheduleId"`
	CabinTypeId int `json:"cabinTypeId"`
}

type GetTicketsRequest struct {
	UserId           *int
	ScheduleId       *int
	BookingReference *string
}

type GetTicketsResponse struct {
	Tickets []Ticket `json:"tickets"`
}
