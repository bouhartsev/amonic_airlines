package domain

type Answers struct {
	Q1 int `json:"q1"`
	Q2 int `json:"q2"`
	Q3 int `json:"q3"`
	Q4 int `json:"q4"`
	Q5 int `json:"q5"`
	Q6 int `json:"q6"`
	Q7 int `json:"q7"`
}

type AddReviewRequest struct {
	From        string  `json:"from"`
	To          string  `json:"to"`
	Age         int     `json:"age"`
	Gender      int     `json:"gender"`
	CabinTypeId int     `json:"cabinTypeId"`
	Answers     Answers `json:"answers"`
}

type GetBriefReviewsRequest struct {
	From string `json:"from"`
	To   string `json:"to"`
}

type Review struct {
	Gender struct {
		Male   int
		Female int
	} `json:"gender"`
	Age struct {
		R1824 int `json:"18-24"`
		R2539 int `json:"25-39"`
		R4059 int `json:"40-59"`
		R60   int `json:"60"`
	} `json:"age"`
	CabinType struct {
		Economy  int `json:"economy"`
		Business int `json:"business"`
		First    int `json:"first"`
	} `json:"cabinType"`
	DestinationAirport struct {
		AUH int `json:"AUH"`
		BAH int `json:"BAH"`
		DOH int `json:"DOH"`
		RYU int `json:"RYU"`
		CAI int `json:"CAI"`
	}
}

type GetBriefReviewsResponse struct {
	Total  int    `json:"total"`
	Review Review `json:"review"`
}

type GetDetailedReviewsRequest struct {
	From   string `json:"-"`
	To     string `json:"-"`
	Age    *int   `json:"-"`
	Gender *int   `json:"-"`
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
	Total  int `json:"total"`
	Gender struct {
		Male   int
		Female int
	} `json:"gender"`
	Age struct {
		R1824 int `json:"18-24"`
		R2539 int `json:"25-39"`
		R4059 int `json:"40-59"`
		R60   int `json:"60"`
	} `json:"age"`
	CabinType struct {
		Economy  int `json:"economy"`
		Business int `json:"business"`
		First    int `json:"first"`
	} `json:"cabinType"`
	DestinationAirport struct {
		AUH int `json:"AUH"`
		BAH int `json:"BAH"`
		DOH int `json:"DOH"`
		RYU int `json:"RYU"`
		CAI int `json:"CAI"`
	}
}
