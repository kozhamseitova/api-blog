package pgrepo

import (
	"api-blog/internal/entity"
	"context"
	"fmt"
)

func (p *Postgres) GetCategories(ctx context.Context) ([]entity.Category, error) {
	query := fmt.Sprintf(
		`SELECT * from %s`, categoriesTable)

	rows, err := p.Pool.Query(ctx, query)
	if err != nil {
		return nil, err
	}

	var categories []entity.Category

	for rows.Next() {
		category := entity.Category{}

		rows.Scan(
			&category.ID,
			&category.Name,
		)

		categories = append(categories, category)
	}

	return categories, nil
}
