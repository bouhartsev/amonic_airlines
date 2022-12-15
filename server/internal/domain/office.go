package domain

type (
	Office struct {
		Id        *int    `json:"id,omitempty" example:"12"`
		CountryId *int    `json:"countryId,omitempty" example:"185"`
		Title     *string `json:"title,omitempty" example:"Abu dhabi"`
		Phone     *string `json:"phone,omitempty" example:"252-224-8525"`
		Contact   *string `json:"contact,omitempty" example:"David Johns"`
	}

	GetOfficesResponse struct {
		Offices []*Office `json:"offices,omitempty"`
	}
)
