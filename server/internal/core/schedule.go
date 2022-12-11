package core

import (
	"context"
	"fmt"
	"strings"

	"github.com/bouhartsev/amonic_airlines/server/internal/domain"
	"github.com/bouhartsev/amonic_airlines/server/internal/domain/errdomain"
)

func (c *Core) GetSchedules(ctx context.Context, request *domain.GetSchedulesRequest) (*domain.GetSchedulesResponse, error) {
	var (
		args  []any
		query []string
	)

	if request.FlightNumber != nil {
		query = append(query, "s.flightNumber = ?")
		args = append(args, *request.FlightNumber)
	}

	if request.Outbound != nil {
		query = append(query, "s.date >= ?")
		args = append(args, *request.Outbound)
	}

	if request.To != nil {
		query = append(query, "a1.Name = ?")
		args = append(args, *request.To)
	}

	if request.From != nil {
		query = append(query, "a2.Name = ?")
		args = append(args, *request.From)
	}

	var sort string
	if request.SortBy != nil {
		switch *request.SortBy {
		case "datetime":
			sort = "date, time"
		case "price":
			sort = "s.economyPrice"
		case "confirmed":
			sort = "s.confirmed DESC"
		case "unconfirmed":
			sort = "s.confirmed ASC"
		default:
			sort = "date, time"
		}
	} else {
		sort = "date, time"
	}

	var where string
	if len(query) > 0 {
		where = "where " + strings.Join(query, " AND ")
	}

	q := fmt.Sprintf(`select s.id as id, s.date as date, s.time as time,
                                    a1.Name as air_to, a2.Name as air_from,
                                    s.FlightNumber as flight_number,
                                    s.confirmed,
                                    air.Name,
                                    s.EconomyPrice as economy_price
                             from schedules s
                                  inner join routes r on s.RouteID = r.ID
                                  inner join airports a1
                                      on r.ArrivalAirportID = a1.ID
                                  inner join airports a2
                                      on r.DepartureAirportID = a2.ID
                                  inner join aircrafts air
                                      on air.ID = s.AircraftID
                             %s
                             order by %s`, where, sort)

	rows, err := c.db.QueryContext(ctx, q, args...)

	if err != nil {
		return nil, errdomain.NewInternalError(err.Error())
	}

	defer rows.Close()

	var schedules []domain.Schedule
	for rows.Next() {
		var s domain.Schedule

		err = rows.Scan(
			&s.Id,
			&s.Date,
			&s.Time,
			&s.To,
			&s.From,
			&s.FlightNumber,
			&s.Confirmed,
			&s.Aircraft,
			&s.EconomyPrice,
		)

		if err != nil {
			return nil, errdomain.NewInternalError(err.Error())
		}

		s.CalculatePrices()
		schedules = append(schedules, s)
	}

	return &domain.GetSchedulesResponse{Schedules: schedules}, nil
}

func (c *Core) UpdateSchedule(ctx context.Context, request *domain.UpdateScheduleRequest) error {
	var (
		query []string
		args  []any
	)

	if request.Time != nil {
		query = append(query, "time = ?")
		args = append(args, *request.Time)
	}

	if request.EconomyPrice != nil {
		query = append(query, "economyPrice = ?")
		args = append(args, *request.EconomyPrice)
	}

	if request.Date != nil {
		query = append(query, "date = ?")
		args = append(args, *request.Date)
	}

	args = append(args, request.ScheduleId)

	q := fmt.Sprintf(`update schedules set %s where id = ?`, query)

	_, err := c.db.ExecContext(ctx, q, args...)
	return err
}

func (c *Core) ConfirmSchedule(ctx context.Context, request *domain.ConfirmScheduleRequest) error {
	_, err := c.db.ExecContext(ctx, `update schedules set confirmed = 1 where id = ?`, request.ScheduleId)
	return err
}

func (c *Core) UnconfirmSchedule(ctx context.Context, request *domain.UnconfirmScheduleRequest) error {
	_, err := c.db.ExecContext(ctx, `update schedules set confirmed = 0 where id = ?`, request.ScheduleId)
	return err
}
