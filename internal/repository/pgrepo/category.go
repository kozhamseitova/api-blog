package pgrepo

import (
	"api-blog/internal/entity"
	"context"
	"fmt"
	"github.com/georgysavva/scany/pgxscan"
)

func (p *Postgres) GetCategories(ctx context.Context) ([]*entity.Category, error) {
	query := fmt.Sprintf(
		`SELECT * from %s`, categoriesTable)

	var categories []*entity.Category

	err := pgxscan.Select(ctx, p.Pool, &categories, query)
	if err != nil {
		return nil, err
	}

	return categories, nil
}
