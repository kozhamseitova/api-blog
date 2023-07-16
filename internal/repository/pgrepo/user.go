package pgrepo

import (
	"api-blog/internal/entity"
	"context"
	"fmt"
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

func (p *Postgres) Login(ctx context.Context, username, password string) (entity.User, error) {
	query := fmt.Sprintf(`
		SELECT * FROM %s WHERE
			username = $1 
			LIMIT 1
	`, usersTable)

	row := p.Pool.QueryRow(ctx, query, username)

	var user entity.User
	err := row.Scan(&user.ID, &user.Username, &user.Password, &user.FirstName, &user.LastName) // Adjust the fields as per your entity.User structure
	return user, err
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
		return err
	}

	return nil
}

func (p *Postgres) DeleteUser(ctx context.Context, id int64) error {
	query := fmt.Sprintf(
		`DELETE FROM %s
				WHERE id = $1`, usersTable)

	_, err := p.Pool.Exec(ctx, query, id)
	if err != nil {
		return err
	}

	return nil
}
