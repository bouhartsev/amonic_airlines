package core

import (
	"context"
	"fmt"
	"time"

	"github.com/bouhartsev/amonic_airlines/server/internal/domain"
)

func (c *Core) GetDetailedReport(ctx context.Context) (*domain.GetDetailedReportResponse, error) {
	now := time.Now()
	var resp domain.GetDetailedReportResponse
	const day = 30

	// Flights
	row := c.db.QueryRowContext(ctx, fmt.Sprintf(`select
    	  sum(if(s.confirmed = 1, 1, 0)) as confirmed,
    	  sum(if(s.confirmed = 0, 1, 0)) as cancelled,
    	  ceiling(avg(r.FlightTime)) as average_daily_flight_time
		  from schedules s, routes r
		  where s.date > date_sub(now(), interval %d day)`, day))
	if err := row.Scan(
		&resp.Flights.Confirmed,
		&resp.Flights.Cancelled,
		&resp.Flights.AverageDailyFlightTimeMinutes,
	); err != nil {
		fmt.Println("flights error", err.Error())
		return nil, err
	}

	// Top Customers
	rows, err := c.db.QueryContext(ctx, fmt.Sprintf(`select count(t.userid) as total, u.FirstName, u.LastName
                                                            from tickets t
                                                            join users u on t.UserID = u.ID
                                                            join schedules s on t.ScheduleID = s.ID
                                                            where t.Confirmed and s.Confirmed and s.Date > date_sub(now(), interval %d day)
                                                            group by u.FirstName, u.LastName
                                                            order by total desc
                                                            limit 3`, day))
	if err != nil {
		fmt.Println("top customers error", err.Error())
		return nil, err
	}
	defer rows.Close()

	var first, second string
	var topCustomers []domain.TopCustomer
	for rows.Next() {
		var c domain.TopCustomer

		if err = rows.Scan(
			&c.TicketNumber,
			&first,
			&second,
		); err != nil {
			fmt.Println("scan top customers error", err.Error())
			return nil, err
		}

		c.Name = fmt.Sprintf("%s %s", first, second)
		topCustomers = append(topCustomers, c)
	}

	resp.TopCustomers.First = topCustomers[0]
	resp.TopCustomers.Second = topCustomers[1]
	resp.TopCustomers.Third = topCustomers[2]

	// number of passengers flying busiest day
	row = c.db.QueryRowContext(ctx, fmt.Sprintf(`select max(tab.total) as max, DATE_FORMAT(tab.Date, '%Y-%m-%d')
                                                        from(select count(*) as total, s.Date
                                                             from tickets t
                                                             join schedules s on t.ScheduleID = s.ID
                                                             where s.Confirmed and t.Confirmed and s.Date > date_sub(now(), interval %d day)
                                                             group by t.ScheduleID) tab
                                                        group by tab.Date
                                                        order by max desc
                                                        limit 1`, day))

	if err := row.Scan(
		&resp.NumberOfPassengersFlying.BusiestDay.FlyingNumber,
		&resp.NumberOfPassengersFlying.BusiestDay.Date,
	); err != nil {
		fmt.Println("NumberOfPassengersFlying busiest day error", err.Error())
		return nil, err
	}

	// number of passengers flying most quiet day
	row = c.db.QueryRowContext(ctx, fmt.Sprintf(`select min(tab.total) as min, DATE_FORMAT(tab.Date, '%Y-%m-%d')
                                                        from(select count(*) as total, s.Date
                                                             from tickets t
                                                             join schedules s on t.ScheduleID = s.ID
                                                             where s.Confirmed and t.Confirmed and s.Date > date_sub(now(), interval %d day)
                                                             group by t.ScheduleID) tab
                                                        group by tab.Date
                                                        order by min
                                                        limit 1;`, day))

	if err := row.Scan(
		&resp.NumberOfPassengersFlying.MostQuietDay.FlyingNumber,
		&resp.NumberOfPassengersFlying.MostQuietDay.Date,
	); err != nil {
		fmt.Println("NumberOfPassengersFlying most quiet day error", err.Error())
		return nil, err
	}

	// Top Offices
	rows, err = c.db.QueryContext(ctx, fmt.Sprintf(`select count(t.userid) as total, o.Title
                                                           from tickets t
                                                           join users u on t.UserID = u.ID
                                                           join offices o on u.OfficeID = o.ID
                                                           join schedules s on t.ScheduleID = s.ID
                                                           where t.Confirmed and s.Confirmed and s.Date > date_sub(now(), interval %d day)
                                                           group by o.Title
                                                           order by total desc
                                                           limit 3`, day))
	if err != nil {
		fmt.Println("top offices error", err.Error())
		return nil, err
	}
	defer rows.Close()

	var topOffices []domain.TopOffice
	for rows.Next() {
		var c domain.TopOffice

		if err = rows.Scan(
			&c.Total,
			&c.Name,
		); err != nil {
			fmt.Println("scan top offices error", err.Error())
			return nil, err
		}
		topOffices = append(topOffices, c)
	}

	resp.TopOffices.First = topOffices[0]
	resp.TopOffices.Second = topOffices[1]
	resp.TopOffices.Third = topOffices[2]

	// ticket sales yesterday
	row = c.db.QueryRowContext(ctx, `select sum(s.EconomyPrice + s.EconomyPrice * 1.35 + s.EconomyPrice * 1.35 * 1.3) as total
                                                        from tickets t
                                                        join schedules s on t.ScheduleID = s.ID
                                                        where t.Confirmed and s.Confirmed and s.Date > date_add(now(), interval -1 day)`)
	if err := row.Scan(
		&resp.TicketSales.Yesterday,
	); err != nil {
		fmt.Println("scan sales yesterday error", err.Error())
		return nil, err
	}

	// ticket sales 2 days ago
	row = c.db.QueryRowContext(ctx, `select sum(s.EconomyPrice + s.EconomyPrice * 1.35 + s.EconomyPrice * 1.35 * 1.3) as total
                                                        from tickets t
                                                        join schedules s on t.ScheduleID = s.ID
                                                        where t.Confirmed and s.Confirmed and s.Date > date_add(now(), interval -2 day)`)
	if err := row.Scan(
		&resp.TicketSales.TwoDaysAgo,
	); err != nil {
		fmt.Println("scan sales 2d error", err.Error())
		return nil, err
	}

	// ticket sales 3 days ago
	row = c.db.QueryRowContext(ctx, `select sum(s.EconomyPrice + s.EconomyPrice * 1.35 + s.EconomyPrice * 1.35 * 1.3) as total
                                                        from tickets t
                                                        join schedules s on t.ScheduleID = s.ID
                                                        where t.Confirmed and s.Confirmed and s.Date > date_add(now(), interval -3 day)`)
	if err := row.Scan(
		&resp.TicketSales.ThreeDaysAgo,
	); err != nil {
		fmt.Println("scan sales 3d error", err.Error())
		return nil, err
	}

	// empty seats this week
	row = c.db.QueryRowContext(ctx, `select sum(a.TotalSeats - tab.total_tickets)
                                                            from (select count(*) as total_tickets, t.ScheduleID, s.AircraftID
                                                        from tickets t
                                                        join schedules s on t.ScheduleID = s.ID
                                                        join aircrafts a on s.AircraftID = a.ID
                                                        where s.Date > date_add(now(), interval -7 day)
                                                        group by t.ScheduleID) tab
                                                        join aircrafts a on a.id = tab.AircraftID`)

	if err := row.Scan(
		&resp.EmptySeats.ThisWeek,
	); err != nil {
		fmt.Println("scan empty seats 1w error", err.Error())
		return nil, err
	}

	// empty seats two weeks
	row = c.db.QueryRowContext(ctx, `select sum(a.TotalSeats - tab.total_tickets)
                                                            from (select count(*) as total_tickets, t.ScheduleID, s.AircraftID
                                                        from tickets t
                                                        join schedules s on t.ScheduleID = s.ID
                                                        join aircrafts a on s.AircraftID = a.ID
                                                        where s.Date > date_add(now(), interval -14 day)
                                                        group by t.ScheduleID) tab
                                                        join aircrafts a on a.id = tab.AircraftID`)

	if err := row.Scan(
		&resp.EmptySeats.LastWeek,
	); err != nil {
		fmt.Println("scan empty seats 2w error", err.Error())
		return nil, err
	}

	// empty seats two weeks
	row = c.db.QueryRowContext(ctx, `select sum(a.TotalSeats - tab.total_tickets)
                                                            from (select count(*) as total_tickets, t.ScheduleID, s.AircraftID
                                                        from tickets t
                                                        join schedules s on t.ScheduleID = s.ID
                                                        join aircrafts a on s.AircraftID = a.ID
                                                        where s.Date > date_add(now(), interval -21 day)
                                                        group by t.ScheduleID) tab
                                                        join aircrafts a on a.id = tab.AircraftID`)

	if err := row.Scan(
		&resp.EmptySeats.TwoWeeksAgo,
	); err != nil {
		fmt.Println("scan empty seats 3w error", err.Error())
		return nil, err
	}

	elapsed := time.Since(now)
	resp.GenerateTime = elapsed.Seconds()

	return &resp, nil
}
