package domain

type Airport struct {
	Id        *int    `json:"id,omitempty"`
	CountryId *int    `json:"countryId,omitempty"`
	Name      *string `json:"name,omitempty"`
	IATACode  *string `json:"IATACode,omitempty"`
}

type GetAirportsResponse struct {
	Airports []Airport `json:"airports"`
}

type CabinType struct {
	Id   *int    `json:"id,omitempty"`
	Name *string `json:"name,omitempty"`
}

type GetCabinTypesResponse struct {
	CabinTypes []CabinType `json:"cabinTypes"`
}

type Country struct {
	Id   *int    `json:"id,omitempty"`
	Name *string `json:"name,omitempty"`
}

type GetCountriesResponse struct {
	Countries []Country `json:"countries"`
}
