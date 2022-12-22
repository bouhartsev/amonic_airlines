package core

import (
	"context"

	"github.com/bouhartsev/amonic_airlines/server/internal/domain"
)

func (c *Core) GetAirports(ctx context.Context) (*domain.GetAirportsResponse, error) {
	rows, err := c.db.QueryContext(ctx, `select id, countryId, IATACode, Name from airports`)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var airports []domain.Airport

	for rows.Next() {
		var a domain.Airport

		err = rows.Scan(
			&a.Id,
			&a.CountryId,
			&a.IATACode,
			&a.Name,
		)

		if err != nil {
			return nil, err
		}

		airports = append(airports, a)
	}

	return &domain.GetAirportsResponse{Airports: airports}, nil
}

func (c *Core) GetCabinTypes(ctx context.Context) (*domain.GetCabinTypesResponse, error) {
	rows, err := c.db.QueryContext(ctx, `select id, name from cabintypes`)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var cbs []domain.CabinType

	for rows.Next() {
		var cb domain.CabinType

		err = rows.Scan(
			&cb.Id,
			&cb.Name,
		)

		if err != nil {
			return nil, err
		}

		cbs = append(cbs, cb)
	}

	return &domain.GetCabinTypesResponse{CabinTypes: cbs}, nil
}

func (c *Core) GetCountries(ctx context.Context) (*domain.GetCountriesResponse, error) {
	rows, err := c.db.QueryContext(ctx, `select id, name from countries`)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var countries []domain.Country

	for rows.Next() {
		var c domain.Country

		err = rows.Scan(
			&c.Id,
			&c.Name,
		)

		if err != nil {
			return nil, err
		}

		countries = append(countries, c)
	}

	return &domain.GetCountriesResponse{Countries: countries}, nil
}
