package core

import (
	"context"
	"fmt"

	"github.com/bouhartsev/amonic_airlines/server/internal/domain"
)

func (c *Core) GetAmenities(ctx context.Context) (*domain.GetAmenitiesResponse, error) {
	var am []domain.Amenity

	rows, err := c.db.QueryContext(ctx, `select id, service, price from amenities`)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	for rows.Next() {
		var a domain.Amenity

		err = rows.Scan(
			&a.Id,
			&a.Description,
			&a.Price,
		)
		if err != nil {
			return nil, err
		}

		am = append(am, a)
	}

	return &domain.GetAmenitiesResponse{Amenities: am}, nil
}

func (c *Core) GetTicketAmenities(ctx context.Context, ticketId int) (*domain.GetTicketAmenitiesResponse, error) {
	var am []domain.TicketAmenity

	rows, err := c.db.QueryContext(ctx, `select a.id, a.price from amenitiestickets as att join amenities as a on a.id = att.amenityId where att.ticketId = ?`, ticketId)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	for rows.Next() {
		var a domain.TicketAmenity

		err = rows.Scan(
			&a.AmenityId,
			&a.Price,
		)
		if err != nil {
			return nil, err
		}

		am = append(am, a)
	}

	return &domain.GetTicketAmenitiesResponse{Amenities: am}, nil
}

func (c *Core) RemoveTicketAmenities(ctx context.Context, req *domain.RemoveTicketAmenitiesRequest) error {
	for _, a := range req.Amenities {
		_, err := c.db.ExecContext(ctx, `delete from amenitiestickets where ticketId = ? and amenityid = ?`, req.TicketId, a)
		if err != nil {
			return err
		}
	}

	return nil
}

func (c *Core) AddTicketAmenities(ctx context.Context, req *domain.AddTicketAmenitiesRequest) error {
	amenities, err := c.GetAmenities(ctx)
	if err != nil {
		return err
	}

	for _, a := range req.Amenities {
		var amenityId int
		var price float32

		for _, am := range amenities.Amenities {
			if am.Id == a {
				amenityId = am.Id
				price = am.Price
			}
		}

		_, err = c.db.ExecContext(ctx, `insert into amenitiestickets(amenityId, ticketId, price) values(?, ?, ?)`, amenityId, req.TicketId, price)
		if err != nil {
			return err
		}
	}

	return nil
}

func (c *Core) GetAmenitiesBriefReport(ctx context.Context, req *domain.GetAmenitiesBriefReportRequest) (*domain.GetAmenitiesBriefReportResponse, error) {
	where := ""
	args := make([]any, 0)
	args = append(args, 1)

	if req.ScheduleId != nil {
		where += fmt.Sprintf("and t.scheduleId = ?")
		args = append(args, *req.ScheduleId)
	}
	if req.DateFrom != nil {
		where += fmt.Sprintf("and s.date >= ?")
		args = append(args, *req.DateFrom)
	}
	if req.DateTo != nil {
		where += fmt.Sprintf("and s.date <= ?")
		args = append(args, *req.DateTo)
	}

	q := fmt.Sprintf(`select count(*), a.id, a.service
		  from amenities a
		  join amenitiestickets at on at.AmenityID = a.id
		  join tickets t on at.TicketID = t.ID
		  join schedules s on t.ScheduleID = s.ID
		  where t.CabinTypeID = ? %s
		  group by a.id, a.service, a.price`, where)

	resp := &domain.GetAmenitiesBriefReportResponse{}

	// Economy
	rows, err := c.db.QueryContext(ctx, q, args...)
	if err != nil {
		return nil, err
	}

	defer rows.Close()
	var reports []domain.AmenityReport

	for rows.Next() {
		var r domain.AmenityReport

		err = rows.Scan(
			&r.Count,
			&r.AmenityId,
			&r.Description,
		)
		if err != nil {
			return nil, err
		}

		reports = append(reports, r)
	}

	resp.Economy.Reports = reports

	// Business
	args[0] = 2
	rows, err = c.db.QueryContext(ctx, q, args...)
	if err != nil {
		return nil, err
	}

	reports = make([]domain.AmenityReport, 0)

	for rows.Next() {
		var r domain.AmenityReport

		err = rows.Scan(
			&r.Count,
			&r.AmenityId,
			&r.Description,
		)
		if err != nil {
			return nil, err
		}

		reports = append(reports, r)
	}

	resp.Business.Reports = reports

	// First Class
	args[0] = 3
	rows, err = c.db.QueryContext(ctx, q, args...)
	if err != nil {
		return nil, err
	}

	reports = make([]domain.AmenityReport, 0)

	for rows.Next() {
		var r domain.AmenityReport

		err = rows.Scan(
			&r.Count,
			&r.AmenityId,
			&r.Description,
		)
		if err != nil {
			return nil, err
		}

		reports = append(reports, r)
	}

	resp.FirstClass.Reports = reports

	return resp, nil
}
