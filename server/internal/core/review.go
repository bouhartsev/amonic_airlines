package core

import (
	"context"
	"database/sql"
	"errors"

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
	return nil, nil
}

func getDetailedReview(ctx context.Context, req *domain.GetDetailedReviewsRequest) (*domain.DetailedReview, error) {

	return nil, nil
}
