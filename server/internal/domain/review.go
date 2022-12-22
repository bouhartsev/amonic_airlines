package domain

type Answers struct {
	Q1 int `json:"q1" example:"5"`
	Q2 int `json:"q2" example:"4"`
	Q3 int `json:"q3" example:"1"`
	Q4 int `json:"q4" example:"2"`
	Q5 int `json:"q5" example:"3"`
	Q6 int `json:"q6" example:"7"`
	Q7 int `json:"q7" example:"6"`
}

type AddReviewRequest struct {
	From        string  `json:"from" example:"AUH"`
	To          string  `json:"to" example:"DOH"`
	Age         int     `json:"age" example:"20"`
	Gender      int     `json:"gender" example:"1"`
	CabinTypeId int     `json:"cabinTypeId" example:"2"`
	Answers     Answers `json:"answers"`
}

type GetBriefReviewsRequest struct {
	From string `json:"-"`
	To   string `json:"-"`
}

type Review struct {
	Gender struct {
		Male   int `json:"male" example:"110"`
		Female int `json:"female" example:"24"`
	} `json:"gender"`
	Age struct {
		R1824 int `json:"18-24" example:"24"`
		R2539 int `json:"25-39" example:"129"`
		R4059 int `json:"40-59" example:"32"`
		R60   int `json:"60" example:"75"`
	} `json:"age"`
	CabinType struct {
		Economy  int `json:"economy" example:"23"`
		Business int `json:"business" example:"12"`
		First    int `json:"first" example:"88"`
	} `json:"cabinType"`
	DestinationAirport struct {
		AUH int `json:"AUH" example:"55"`
		BAH int `json:"BAH" example:"32"`
		DOH int `json:"DOH" example:"12"`
		RYU int `json:"RYU" example:"67"`
		CAI int `json:"CAI" example:"109"`
	}
}

type GetBriefReviewsResponse struct {
	Total  int    `json:"total" example:"1132"`
	Review Review `json:"review"`
}

type GetDetailedReviewsRequest struct {
	From string `json:"-"`
	To   string `json:"-"`
}

type GetDetailedReviewsResponse struct {
	Q1 DetailedReview `json:"q1"`
	Q2 DetailedReview `json:"q2"`
	Q3 DetailedReview `json:"q3"`
	Q4 DetailedReview `json:"q4"`
}

type DetailedReview struct {
	R1 DetailedReviewSub `json:"1"`
	R2 DetailedReviewSub `json:"2"`
	R3 DetailedReviewSub `json:"3"`
	R4 DetailedReviewSub `json:"4"`
	R5 DetailedReviewSub `json:"5"`
	R6 DetailedReviewSub `json:"6"`
	R7 DetailedReviewSub `json:"7"`
}

type DetailedReviewSub struct {
	Total  int `json:"total" example:"3086"`
	Gender struct {
		Male   int `json:"male" example:"110"`
		Female int `json:"female" example:"24"`
	} `json:"gender"`
	Age struct {
		R1824 int `json:"18-24" example:"24"`
		R2539 int `json:"25-39" example:"129"`
		R4059 int `json:"40-59" example:"32"`
		R60   int `json:"60" example:"75"`
	} `json:"age"`
	CabinType struct {
		Economy  int `json:"economy" example:"23"`
		Business int `json:"business" example:"12"`
		First    int `json:"first" example:"88"`
	} `json:"cabinType"`
	DestinationAirport struct {
		AUH int `json:"AUH" example:"55"`
		BAH int `json:"BAH" example:"32"`
		DOH int `json:"DOH" example:"12"`
		RYU int `json:"RYU" example:"67"`
		CAI int `json:"CAI" example:"109"`
	}
}
