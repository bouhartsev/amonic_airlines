package core

import (
	"context"

	"github.com/bouhartsev/amonic_airlines/server/internal/domain"
	"github.com/bouhartsev/amonic_airlines/server/internal/domain/errdomain"
)

func (c *Core) GetOffices(ctx context.Context) (*domain.GetOfficesResponse, error) {
	rows, err := c.db.QueryContext(ctx, `select id, countryId, title, phone, contact from offices`)

	if err != nil {
		return nil, errdomain.NewInternalError(err.Error())
	}

	var offices []*domain.Office

	for rows.Next() {
		var office domain.Office

		err = rows.Scan(
			&office.Id,
			&office.CountryId,
			&office.Title,
			&office.Phone,
			&office.Contact,
		)

		if err != nil {
			c.logger.Error(err.Error())
		}

		offices = append(offices, &office)
	}

	return &domain.GetOfficesResponse{Offices: offices}, nil
}
