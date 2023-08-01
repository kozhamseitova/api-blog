package pgrepo

import (
	"github.com/jackc/pgx/v4/pgxpool"
)

const (
	usersTable              = "users"
	articlesTable           = "articles"
	articlesCategoriesTable = "articles_categories"
	categoriesTable         = "categories"
)

type Postgres struct {
	Pool *pgxpool.Pool
}

func New(pool *pgxpool.Pool) *Postgres {
	return &Postgres{Pool: pool}
}
