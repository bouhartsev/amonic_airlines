package core

import (
	"context"
	"math/rand"

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
