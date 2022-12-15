package domain

type Airport struct {
	Id        *int    `json:"id,omitempty" example:"10"`
	CountryId *int    `json:"countryId,omitempty" example:"123"`
	Name      *string `json:"name,omitempty" example:"Abu Dhabi"`
	IATACode  *string `json:"IATACode,omitempty" example:"AUH"`
}

type GetAirportsResponse struct {
	Airports []Airport `json:"airports"`
}

type CabinType struct {
	Id   *int    `json:"id,omitempty" example:"3"`
	Name *string `json:"name,omitempty" example:"First Class"`
}

type GetCabinTypesResponse struct {
	CabinTypes []CabinType `json:"cabinTypes"`
}

type Country struct {
	Id   *int    `json:"id,omitempty" example:"13"`
	Name *string `json:"name,omitempty" example:"Armenia"`
}

type GetCountriesResponse struct {
	Countries []Country `json:"countries"`
}
