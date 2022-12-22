package core

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"io"
	"mime/multipart"
	"regexp"
	"strconv"
	"strings"
	"sync"

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

	q := fmt.Sprintf(`select s.id as id, s.date as schedule_date, s.time as schedule_time,
                                    a1.Name as air_to, a2.Name as air_from,
                                    s.FlightNumber as flight_number,
                                    s.confirmed as confirmed,
                                    air.Name as aircraft_name,
                                    s.EconomyPrice as economy_price,
                                    air.TotalSeats - count(t.ID) as empty_seats
                             from schedules s
                                      inner join routes r on s.RouteID = r.ID
                                      inner join airports a1
                                                 on r.ArrivalAirportID = a1.ID
                                      inner join airports a2
                                                 on r.DepartureAirportID = a2.ID
                                      inner join aircrafts air
                                                 on air.ID = s.AircraftID
                                      inner join tickets t
                                                 on s.ID = t.ScheduleID
                             %s
                             group by id, schedule_date, schedule_time, air_from, flight_number, confirmed, aircraft_name, economy_price
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
			&s.EmptySeats,
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

func (c *Core) UpdateSchedulesFromFile(ctx context.Context, file multipart.File) (*domain.UpdateSchedulesFromFileResponse, error) {
	content, err := io.ReadAll(file)
	if err != nil {
		return nil, err
	}

	fileRows := strings.Split(string(content), "\r\n")
	var (
		success       int
		duplicates    int
		missingFields int
		actions       = make([]*domain.ScheduleAction, 0)
	)

	actionList, duplicates := removeScheduleActionDuplicates(fileRows)

	for _, row := range actionList {
		parts := strings.Split(strings.TrimSpace(row), ",")
		if len(parts) != 9 {
			missingFields++
			continue
		}

		action, ok := parseScheduleActionFromString(parts)
		if ok {
			actions = append(actions, action)
		}
	}

	var (
		wg            sync.WaitGroup
		susccessMutex = sync.Mutex{}
	)
	wg.Add(len(actions))
	for _, a := range actions {
		go func(action *domain.ScheduleAction) {
			defer wg.Done()

			var scheduleId int
			err := c.db.QueryRowContext(ctx, "select id from schedules where flightNumber = ? and date = ?", a.FlightNumber, a.Date).Scan(&scheduleId)
			if err != nil {
				if errors.Is(err, sql.ErrNoRows) {
					return
				}

				c.logger.Error(err.Error())
				return
			}

			var confirmed bool
			if action.Confirmed == "OK" {
				confirmed = true
			} else {
				confirmed = false
			}

			if action.Action == "ADD" {
				q := "insert into schedules(date, time, aircraftId, economyPrice, flightNumber, confirmed, routeId) values(?, ?, ?, ?, ?, ?, (select id from routes where DepartureAirportID = (select id from airports where IATACode = ?) and ArrivalAirportID = (select id from airports where IATACode = ?)))"
				_, err = c.db.ExecContext(ctx, q, action.Date, action.Time, action.Aircraft, action.Price, action.FlightNumber, confirmed, action.From, action.To)
				if err != nil {
					fmt.Println("insert error", action)
					c.logger.Error(err.Error())
					return
				}
				susccessMutex.Lock()
				success++
				susccessMutex.Unlock()
			} else {
				q := "update schedules set time = ?, aircraftId = ?, economyPrice = ?, confirmed = ?, routeId = (select id from routes where DepartureAirportID = (select id from airports where IATACode = ?) and ArrivalAirportID = (select id from airports where IATACode = ?)) where date = ? and flightNumber = ?"
				_, err = c.db.ExecContext(ctx, q, action.Time, action.Aircraft, action.Price, confirmed, action.From, action.To, action.Date, action.FlightNumber)
				if err != nil {
					fmt.Println("update error", action)
					c.logger.Error(err.Error())
					return
				}
				susccessMutex.Lock()
				success++
				susccessMutex.Unlock()
			}
		}(a)
	}
	wg.Wait()

	response := &domain.UpdateSchedulesFromFileResponse{
		SuccessfulChangesApplied:         success,
		DuplicateRecordsDiscarded:        duplicates,
		RecordWithMissingFieldsDiscarded: missingFields,
	}

	return response, nil
}

func parseScheduleActionFromString(parts []string) (*domain.ScheduleAction, bool) {
	var a domain.ScheduleAction
	a.Action = parts[0]
	a.Date = parts[1]
	a.Time = parts[2]
	a.FlightNumber = parts[3]
	a.From = parts[4]
	a.To = parts[5]
	a.Aircraft = parts[6]
	a.Price = parts[7]
	a.Confirmed = parts[8]

	if a.Action != "ADD" && a.Action != "EDIT" {
		return nil, false
	}

	if !regexp.MustCompile(`[0-9]{4}-[0-9]{2}-[0-9]{2}`).MatchString(a.Date) {
		return nil, false
	}

	if !regexp.MustCompile(`[0-9]{2}:[0-9]{2}`).MatchString(a.Time) {
		return nil, false
	}

	if !regexp.MustCompile(`[A-Z]{2}`).MatchString(a.From) {
		return nil, false
	}

	if !regexp.MustCompile(`[A-Z]{2}`).MatchString(a.To) {
		return nil, false
	}

	if _, err := strconv.Atoi(a.FlightNumber); err != nil {
		return nil, false
	}

	if _, err := strconv.Atoi(a.Aircraft); err != nil {
		return nil, false
	}

	if _, err := strconv.ParseFloat(a.Price, 32); err != nil {
		return nil, false
	}

	if a.Confirmed != "OK" && a.Confirmed != "CANCELED" {
		return nil, false
	}

	return &a, true
}

func removeScheduleActionDuplicates(rows []string) ([]string, int) {
	var duplicates int

	allKeys := make(map[string]bool)
	var list []string
	for _, item := range rows {
		if _, value := allKeys[item]; !value {
			allKeys[item] = true
		} else {
			delete(allKeys, item)
			duplicates++
		}
	}

	return list, duplicates
}
