package domain

type Ticket struct {
	Id                *int    `json:"id,omitempty" example:"30"`
	UserId            *int    `json:"userId,omitempty" example:"22"`
	ScheduleId        *int    `json:"scheduleId,omitempty" example:"124"`
	CabinTypeId       *int    `json:"cabinTypeId,omitempty" example:"3"`
	FirstName         *string `json:"firstName,omitempty" example:"Pines"`
	LastName          *string `json:"lastName,omitempty" example:"Herbarien"`
	Phone             *string `json:"phone,omitempty" example:"650-995-8364"`
	PassportNumber    *string `json:"passportNumber,omitempty" example:"152442037"`
	PassportCountryId *int    `json:"passportCountryId,omitempty" example:"104"`
	BookingReference  *string `json:"bookingReference,omitempty" example:"AAGERS"`
	Confirmed         *bool   `json:"confirmed,omitempty"`
}

type AddTicketRequest struct {
	Outbound  TicketInfo  `json:"outbound"`
	Return    *TicketInfo `json:"return,omitempty"`
	Passenger struct {
		Firstname         string `json:"firstname" example:"Alex"`
		Lastname          string `json:"lastname" example:"Herbarien"`
		Birthdate         string `json:"birthdate" example:"2017-12-04 17:00:00"`
		PassportNumber    string `json:"passportNumber" example:"152442037"`
		PassportCountryId int    `json:"passportCountryId" example:"12"`
		Phone             string `json:"phone" example:"865-951-7895"`
	}
}

type TicketInfo struct {
	ScheduleId  int `json:"scheduleId" example:"123"`
	CabinTypeId int `json:"cabinTypeId" example:"321"`
}

type GetTicketsRequest struct {
	UserId           *int
	ScheduleId       *int
	BookingReference *string
}

type GetTicketsResponse struct {
	Tickets []Ticket `json:"tickets"`
}
