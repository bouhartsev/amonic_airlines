package core

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

	"github.com/bouhartsev/amonic_airlines/server/internal/domain"
)

func (c *Core) AddReview(ctx context.Context, request *domain.AddReviewRequest) error {
	q := "insert into reviews(`from`, `to`, age, gender, cabinTypeId, q1, q2, q3, q4, q5, q6, q7)" +
		"values(?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)"

	_, err := c.db.ExecContext(ctx, q,
		request.From,
		request.To,
		request.Age,
		request.Gender,
		request.CabinTypeId,
		request.Answers.Q1,
		request.Answers.Q2,
		request.Answers.Q3,
		request.Answers.Q4,
		request.Answers.Q5,
		request.Answers.Q6,
		request.Answers.Q7,
	)

	return err
}

func (c *Core) GetReviewsBrief(ctx context.Context, request *domain.GetBriefReviewsRequest) (*domain.GetBriefReviewsResponse, error) {
	response := domain.GetBriefReviewsResponse{}

	q := `select count(*) from reviews where createdAt >= ? and createdAt <= ?`

	err := c.db.QueryRowContext(ctx, q, request.From, request.To).Scan(&response.Total)
	if err != nil && !errors.Is(err, sql.ErrNoRows) {
		return nil, err
	}

	q = `select count(*) from reviews where gender = 0 and createdAt >= ? and createdAt <= ?`
	err = c.db.QueryRowContext(ctx, q, request.From, request.To).Scan(&response.Review.Gender.Male)
	if err != nil && !errors.Is(err, sql.ErrNoRows) {
		return nil, err
	}

	q = `select count(*) from reviews where gender = 1 and createdAt >= ? and createdAt <= ?`
	err = c.db.QueryRowContext(ctx, q, request.From, request.To).Scan(&response.Review.Gender.Female)
	if err != nil && !errors.Is(err, sql.ErrNoRows) {
		return nil, err
	}

	q = `select count(*) from reviews where age >= 18 AND age <= 24 and createdAt >= ? and createdAt <= ?`
	err = c.db.QueryRowContext(ctx, q, request.From, request.To).Scan(&response.Review.Age.R1824)
	if err != nil && !errors.Is(err, sql.ErrNoRows) {
		return nil, err
	}

	q = `select count(*) from reviews where age >= 25 AND age <= 39 and createdAt >= ? and createdAt <= ?`
	err = c.db.QueryRowContext(ctx, q, request.From, request.To).Scan(&response.Review.Age.R2539)
	if err != nil && !errors.Is(err, sql.ErrNoRows) {
		return nil, err
	}

	q = `select count(*) from reviews where age >= 40 AND age <= 59 and createdAt >= ? and createdAt <= ?`
	err = c.db.QueryRowContext(ctx, q, request.From, request.To).Scan(&response.Review.Age.R4059)
	if err != nil && !errors.Is(err, sql.ErrNoRows) {
		return nil, err
	}

	q = `select count(*) from reviews where age >= 60 and createdAt >= ? and createdAt <= ?`
	err = c.db.QueryRowContext(ctx, q, request.From, request.To).Scan(&response.Review.Age.R60)
	if err != nil && !errors.Is(err, sql.ErrNoRows) {
		return nil, err
	}

	q = `select count(CabinTypeId) from reviews where cabinTypeId = 1 and createdAt >= ? and createdAt <= ? group by cabinTypeId`
	err = c.db.QueryRowContext(ctx, q, request.From, request.To).Scan(
		&response.Review.CabinType.Economy,
	)
	if err != nil && !errors.Is(err, sql.ErrNoRows) {
		return nil, err
	}

	q = `select count(CabinTypeId) from reviews where cabinTypeId = 2 and createdAt >= ? and createdAt <= ? group by cabinTypeId`
	err = c.db.QueryRowContext(ctx, q, request.From, request.To).Scan(
		&response.Review.CabinType.Business,
	)
	if err != nil && !errors.Is(err, sql.ErrNoRows) {
		return nil, err
	}

	q = `select count(CabinTypeId) from reviews where cabinTypeId = 3 and createdAt >= ? and createdAt <= ? group by cabinTypeId`
	err = c.db.QueryRowContext(ctx, q, request.From, request.To).Scan(
		&response.Review.CabinType.Business,
	)
	if err != nil && !errors.Is(err, sql.ErrNoRows) {
		return nil, err
	}

	q = "select count(`to`) from reviews where lower(`to`) = 'AUH' and createdAt >= ? and createdAt <= ? "
	err = c.db.QueryRowContext(ctx, q, request.From, request.To).Scan(
		&response.Review.DestinationAirport.AUH,
	)
	if err != nil && !errors.Is(err, sql.ErrNoRows) {
		return nil, err
	}

	q = "select count(`to`) from reviews where lower(`to`) = 'BAH' and createdAt >= ? and createdAt <= ? "
	err = c.db.QueryRowContext(ctx, q, request.From, request.To).Scan(
		&response.Review.DestinationAirport.BAH,
	)
	if err != nil && !errors.Is(err, sql.ErrNoRows) {
		return nil, err
	}

	q = "select count(`to`) from reviews where lower(`to`) = 'CAI' and createdAt >= ? and createdAt <= ? "
	err = c.db.QueryRowContext(ctx, q, request.From, request.To).Scan(
		&response.Review.DestinationAirport.CAI,
	)
	if err != nil && !errors.Is(err, sql.ErrNoRows) {
		return nil, err
	}

	q = "select count(`to`) from reviews where lower(`to`) = 'DOH' and createdAt >= ? and createdAt <= ? "
	err = c.db.QueryRowContext(ctx, q, request.From, request.To).Scan(
		&response.Review.DestinationAirport.DOH,
	)
	if err != nil && !errors.Is(err, sql.ErrNoRows) {
		return nil, err
	}

	q = "select count(`to`) from reviews where lower(`to`) = 'RYU' and createdAt >= ? and createdAt <= ? "
	err = c.db.QueryRowContext(ctx, q, request.From, request.To).Scan(
		&response.Review.DestinationAirport.RYU,
	)
	if err != nil && !errors.Is(err, sql.ErrNoRows) {
		return nil, err
	}

	return &response, nil
}

func (c *Core) GetDetailedReviews(ctx context.Context, request *domain.GetDetailedReviewsRequest) (*domain.GetDetailedReviewsResponse, error) {
	q1, err := c.getDetailedReviews(ctx, request, "q1")
	if err != nil {
		return nil, err
	}
	q2, err := c.getDetailedReviews(ctx, request, "q2")
	if err != nil {
		return nil, err
	}
	q3, err := c.getDetailedReviews(ctx, request, "q3")
	if err != nil {
		return nil, err
	}
	q4, err := c.getDetailedReviews(ctx, request, "q4")
	if err != nil {
		return nil, err
	}

	return &domain.GetDetailedReviewsResponse{
		Q1: *q1,
		Q2: *q2,
		Q3: *q3,
		Q4: *q4,
	}, nil
}

func (c *Core) getDetailedReviews(ctx context.Context, req *domain.GetDetailedReviewsRequest, quest string) (*domain.DetailedReview, error) {
	q := fmt.Sprintf(`select
    							count(*) as total,
    							sum(if(gender = 0, 1, 0)) as male,
    							sum(if(gender = 1, 1, 0)) as female,
    							sum(if(age >= 18 and age <= 24, 1, 0)) as a1824,
    							sum(if(age >= 25 and age <= 39, 1, 0)) as a2539,
    							sum(if(age >= 40 and age <= 59, 1, 0)) as a4059,
    							sum(if(age >= 60, 1, 0)) as a60,
    							sum(if(CabinTypeId = 1, 1, 0)) as economy,
    							sum(if(CabinTypeId = 2, 1, 0)) as business,
    							sum(if(CabinTypeId = 3, 1, 0)) as first_class,
    							sum(if(lower(`+"`AirportToId`"+`) = 'auh', 1, 0)) as auh,
    							sum(if(lower(`+"`AirportToId`"+`) = 'bah', 1, 0)) as bah,
    							sum(if(lower(`+"`AirportToId`"+`) = 'doh', 1, 0)) as doh,
    							sum(if(lower(`+"`AirportToId`"+`) = 'ryu', 1, 0)) as ryu,
    							sum(if(lower(`+"`AirportToId`"+`) = 'cai', 1, 0)) as cai,
								%s
	                         from reviews
                             where createdAt >= ? and createdAt <= ? `, quest)

	args := []any{req.From, req.To}

	q += fmt.Sprintf(`group by %s order by %s`, quest, quest)

	r := &domain.DetailedReview{}
	reviews := make([]domain.DetailedReviewSub, 0, 7)

	rows, err := c.db.QueryContext(ctx, q, args...)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	for rows.Next() {
		rev := domain.DetailedReviewSub{}
		var question string

		err = rows.Scan(
			&rev.Total,
			&rev.Gender.Male,
			&rev.Gender.Female,
			&rev.Age.R1824,
			&rev.Age.R2539,
			&rev.Age.R4059,
			&rev.Age.R60,
			&rev.CabinType.Economy,
			&rev.CabinType.Business,
			&rev.CabinType.First,
			&rev.DestinationAirport.AUH,
			&rev.DestinationAirport.BAH,
			&rev.DestinationAirport.DOH,
			&rev.DestinationAirport.RYU,
			&rev.DestinationAirport.CAI,
			&question,
		)

		if err != nil {
			return nil, err
		}

		reviews = append(reviews, rev)
	}

	r.R1 = reviews[0]
	r.R2 = reviews[1]
	r.R3 = reviews[2]
	r.R4 = reviews[3]
	r.R5 = reviews[4]
	r.R6 = reviews[5]
	r.R7 = reviews[6]

	return r, nil
}
