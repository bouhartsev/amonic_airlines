package domain

type TopCustomer struct {
	Name         string `json:"name" example:"John Doe"`
	TicketNumber int    `json:"ticketNumber" example:"12"`
}

type TopOffice struct {
	Total int    `json:"total"`
	Name  string `json:"name"`
}

type NumberOfPassengersFlyingDay struct {
	Date         string `json:"date" example:"2022-12-16"`
	FlyingNumber int    `json:"flyingNumber" example:"10"`
}

type GetDetailedReportResponse struct {
	Flights struct {
		Confirmed                     int `json:"confirmed" example:"320"`
		Cancelled                     int `json:"cancelled" example:"160"`
		AverageDailyFlightTimeMinutes int `json:"averageDailyFlightTimeMinutes" example:"68"`
	} `json:"flights"`
	TopCustomers struct {
		First  TopCustomer `json:"first"`
		Second TopCustomer `json:"second"`
		Third  TopCustomer `json:"third"`
	} `json:"topCustomers"`
	NumberOfPassengersFlying struct {
		BusiestDay   NumberOfPassengersFlyingDay `json:"busiestDay"`
		MostQuietDay NumberOfPassengersFlyingDay `json:"mostQuietDay"`
	} `json:"numberOfPassengersFlying"`
	TopOffices struct {
		First  TopOffice `json:"first"`
		Second TopOffice `json:"second"`
		Third  TopOffice `json:"third"`
	} `json:"topOffices"`
	TicketSales struct {
		Yesterday    float32 `json:"yesterday" example:"105.17"`
		TwoDaysAgo   float32 `json:"twoDaysAgo" example:"150.22"`
		ThreeDaysAgo float32 `json:"threeDaysAgo" example:"170.23"`
	} `json:"ticketSales"`
	EmptySeats struct {
		ThisWeek    int `json:"thisWeek" example:"11"`
		LastWeek    int `json:"lastWeek" example:"15"`
		TwoWeeksAgo int `json:"twoWeeksAge" example:"24"`
	} `json:"emptySeats"`
	GenerateTime float64 `json:"generateTimeSeconds" example:"2.13"`
}
