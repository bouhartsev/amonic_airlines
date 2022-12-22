package domain

type Amenity struct {
	Id          int     `json:"id" example:"123"`
	Description string  `json:"description" example:"Next Seat Free"`
	Price       float32 `json:"price" example:"30.0000"`
}

type GetAmenitiesResponse struct {
	Amenities []Amenity `json:"amenities"`
}

type TicketAmenity struct {
	AmenityId int     `json:"amenityId" example:"13"`
	Price     float32 `json:"price" example:"12.0000"`
}

type GetTicketAmenitiesResponse struct {
	Amenities []TicketAmenity `json:"amenities"`
}

type RemoveTicketAmenitiesRequest struct {
	Amenities []int `json:"amenities" example:"1,110,12345,33,12"`
	TicketId  int   `json:"-"`
}

type AddTicketAmenitiesRequest struct {
	Amenities []int `json:"amenities" example:"1,110,12345,33,12"`
	TicketId  int   `json:"-"`
}

type GetAmenitiesBriefReportRequest struct {
	DateFrom   *string `json:"-"`
	DateTo     *string `json:"-"`
	ScheduleId *int    `json:"-"`
}

type AmenityReport struct {
	AmenityId   int    `json:"amenityId" example:"13"`
	Description string `json:"description" example:"Wi-fi 50mb"`
	Count       int    `json:"count" example:"60"`
}

type AmenityBriefReport struct {
	Reports []AmenityReport `json:"reports"`
}

type GetAmenitiesBriefReportResponse struct {
	Economy    AmenityBriefReport `json:"economy"`
	Business   AmenityBriefReport `json:"business"`
	FirstClass AmenityBriefReport `json:"firstClass"`
}
