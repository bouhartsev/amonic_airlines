package domain

type (
	Office struct {
		Id        *int    `json:"id,omitempty"`
		CountryId *int    `json:"countryId,omitempty"`
		Title     *string `json:"title,omitempty"`
		Phone     *string `json:"phone,omitempty"`
		Contact   *string `json:"contact,omitempty"`
	}

	GetOfficesResponse struct {
		Offices []*Office `json:"offices,omitempty"`
	}
)
