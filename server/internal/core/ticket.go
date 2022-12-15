package core

import (
	"context"
	"math/rand"
	"strings"

	"github.com/bouhartsev/amonic_airlines/server/internal/domain"
)

var letters = []rune("ABCDEFGHIJKLMNOPQRSTUVWXYZ123456790")

func (c *Core) AddTicket(ctx context.Context, request *domain.AddTicketRequest) error {
	token, err := c.GetTokenFromContext(ctx)

	if err != nil {
		return err
	}

	bookingRef := make([]rune, 6)
	for i := range bookingRef {
		bookingRef[i] = letters[rand.Intn(len(letters))]
	}

	q := `insert into tickets(userid, scheduleid, cabintypeid, firstname, lastname, phone, passportnumber, passportcountryid, bookingreference, 0)
          values(?,?,?,?,?,?,?,?,?)`

	_, err = c.db.ExecContext(ctx, q,
		token.User.Id,
		request.Outbound.ScheduleId,
		request.Outbound.CabinTypeId,
		request.Passenger.Firstname,
		request.Passenger.Lastname,
		request.Passenger.Phone,
		request.Passenger.PassportNumber,
		request.Passenger.PassportCountryId,
		string(bookingRef),
	)

	if err != nil {
		return err
	}

	if request.Return == nil {
		return nil
	}

	_, err = c.db.ExecContext(ctx, q,
		token.User.Id,
		request.Return.ScheduleId,
		request.Return.CabinTypeId,
		request.Passenger.Firstname,
		request.Passenger.Lastname,
		request.Passenger.Phone,
		request.Passenger.PassportNumber,
		request.Passenger.PassportCountryId,
		string(bookingRef),
	)

	return nil
}

func (c *Core) GetTickets(ctx context.Context, req *domain.GetTicketsRequest) (*domain.GetTicketsResponse, error) {
	var (
		args []any
		vals []string
	)

	q := `select id, userId, scheduleId, cabinTypeId, firstname, lastname, phone, passportNumber, passportCountryId, bookingReference, confirmed from tickets `

	if req.UserId != nil || req.ScheduleId != nil || req.BookingReference != nil {
		q += "where "
	}

	if req.UserId != nil {
		vals = append(vals, "userId = ?")
		args = append(args, *req.UserId)
	}
	if req.ScheduleId != nil {
		vals = append(vals, "scheduledId = ? ")
		args = append(args, *req.ScheduleId)
	}
	if req.BookingReference != nil {
		vals = append(vals, "lower(bookingReference) = lower(?) ")
		args = append(args, *req.BookingReference)
	}

	if len(vals) > 0 {
		s := strings.Join(vals, " AND ")
		q += s
	}

	q += " LIMIT 50"

	var tickets []domain.Ticket

	rows, err := c.db.QueryContext(ctx, q, args...)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	for rows.Next() {
		var t domain.Ticket

		err = rows.Scan(
			&t.Id,
			&t.UserId,
			&t.ScheduleId,
			&t.CabinTypeId,
			&t.FirstName,
			&t.LastName,
			&t.Phone,
			&t.PassportNumber,
			&t.PassportCountryId,
			&t.BookingReference,
			&t.Confirmed,
		)

		if err != nil {
			return nil, err
		}

		tickets = append(tickets, t)
	}

	return &domain.GetTicketsResponse{Tickets: tickets}, nil
}
