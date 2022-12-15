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
