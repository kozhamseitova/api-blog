package pgrepo

import (
	"context"
	"fmt"
	"github.com/georgysavva/scany/pgxscan"
	"github.com/kozhamseitova/api-blog/internal/entity"
	"strings"
)

func (p *Postgres) CreateUser(ctx context.Context, u *entity.User) error {
	query := fmt.Sprintf(`
			INSERT INTO %s (
			                username, -- 1 
			                first_name, -- 2
			                last_name, -- 3
			                password -- 4
			                )
			VALUES ($1, $2, $3, $4)
			`, usersTable)

	_, err := p.Pool.Exec(ctx, query, u.Username, u.FirstName, u.LastName, u.Password)
	if err != nil {
		return err
	}

	return nil
}

func (p *Postgres) GetUserByUsernameAndPassword(ctx context.Context, username, password string) (*entity.User, error) {
	user := new(entity.User)

	query := fmt.Sprintf(`
		SELECT * FROM %s WHERE
			username = $1 
			LIMIT 1
	`, usersTable)

	err := pgxscan.Get(ctx, p.Pool, user, query, strings.TrimSpace(username))

	if err != nil {
		return nil, err
	}

	return user, nil
}

func (p *Postgres) UpdateUser(ctx context.Context, u *entity.User) error {
	query := fmt.Sprintf(`
			UPDATE %s SET 
			              username = $1, 
			              first_name = $2, 
			              last_name = $3
			WHERE id = $4`, usersTable)

	_, err := p.Pool.Exec(ctx, query, u.Username, u.FirstName, u.LastName, u.ID)
	if err != nil {
		return fmt.Errorf("update user err: %w", err)
	}

	return nil
}

func (p *Postgres) DeleteUser(ctx context.Context, id int64) error {
	query := fmt.Sprintf(
		`DELETE FROM %s
				WHERE id = $1`, usersTable)

	_, err := p.Pool.Exec(ctx, query, id)
	if err != nil {
		return fmt.Errorf("delete user err: %w", err)
	}

	return nil
}
