package pgrepo

import (
	"api-blog/internal/entity"
	"context"
	"encoding/json"
	"fmt"
)

func (p *Postgres) CreateArticle(ctx context.Context, a *entity.Article) error {
	tx, err := p.Pool.Begin(ctx)
	if err != nil {
		return err
	}

	queryFirst := fmt.Sprintf(`
		INSERT INTO %s (title, description, user_id)
		VALUES ($1, $2, $3)
		RETURNING id`, articlesTable)

	var articleID int64
	err = tx.QueryRow(ctx, queryFirst, a.Title, a.Description, a.UserID).Scan(&articleID)
	if err != nil {
		_ = tx.Rollback(ctx) // Rollback the transaction on error
		return err
	}

	querySecond := fmt.Sprintf(`
		INSERT INTO %s (article_id, category_id)
		VALUES ($1, $2)`, articlesCategoriesTable)

	for _, category := range a.Categories {
		_, err = tx.Exec(ctx, querySecond, articleID, category.ID)
		if err != nil {
			_ = tx.Rollback(ctx) // Rollback the transaction on error
			return err
		}
	}

	err = tx.Commit(ctx)
	if err != nil {
		return err
	}

	return nil
}

func (p *Postgres) UpdateArticle(ctx context.Context, a *entity.Article) error {
	tx, err := p.Pool.Begin(ctx)
	if err != nil {
		return err
	}
	defer tx.Rollback(ctx)

	updateArticleQuery := fmt.Sprintf(`
						UPDATE %s
						SET title = $1, description = $2, user_id = $3
						WHERE id = $4`, articlesTable)

	_, err = tx.Exec(ctx, updateArticleQuery, a.Title, a.Description, a.UserID, a.ID)
	if err != nil {
		return err
	}

	deleteCategoriesQuery := fmt.Sprintf(
		`DELETE FROM %s
		WHERE article_id = $1`, articlesCategoriesTable)

	_, err = tx.Exec(ctx, deleteCategoriesQuery, a.ID)
	if err != nil {
		return err
	}

	insertCategoriesQuery := fmt.Sprintf(
		`INSERT INTO %s (article_id, category_id)
		VALUES ($1, $2)`, articlesCategoriesTable)

	for _, category := range a.Categories {
		_, err = tx.Exec(ctx, insertCategoriesQuery, a.ID, category.ID)
		if err != nil {
			return err
		}
	}

	err = tx.Commit(ctx)
	if err != nil {
		return err
	}

	return nil
}

func (p *Postgres) DeleteArticleByID(ctx context.Context, id int64) error {
	tx, err := p.Pool.Begin(ctx)
	if err != nil {
		return err
	}
	defer tx.Rollback(ctx)

	deleteCategoriesQuery := fmt.Sprintf(
		`DELETE FROM %s
		WHERE article_id = $1`, articlesCategoriesTable)

	_, err = tx.Exec(ctx, deleteCategoriesQuery, id)
	if err != nil {
		return err
	}

	deleteArticleQuery := fmt.Sprintf(
		`DELETE FROM %s
						WHERE id = $1`, articlesTable)

	_, err = tx.Exec(ctx, deleteArticleQuery, id)
	if err != nil {
		return err
	}

	err = tx.Commit(ctx)
	if err != nil {
		return err
	}

	return nil
}

func (p *Postgres) GetArticleByID(ctx context.Context, id int64) (entity.Article, error) {
	var a entity.Article

	query := fmt.Sprintf(
		`SELECT a.id, a.title, a.description, a.user_id, jsonb_agg(jsonb_build_object('id', c.id, 'name', c.name)) as categories from %s a
				inner join %s ac on a.id = ac.article_id
				inner join %s c on c.id = ac.category_id
				where a.id = $1
				group by a.user_id, a.id`, articlesTable, articlesCategoriesTable, categoriesTable)

	row := p.Pool.QueryRow(ctx, query, id)

	var categoriesData []byte

	err := row.Scan(
		&a.ID,
		&a.Title,
		&a.Description,
		&a.UserID,
		&categoriesData,
	)

	if err != nil {
		return a, err
	}

	var categories []*entity.Category
	err = json.Unmarshal(categoriesData, &categories)
	if err != nil {
		return a, err
	}
	a.Categories = make([]entity.Category, len(categories))

	for i, c := range categories {
		a.Categories[i] = *c
	}

	return a, nil
}

func (p *Postgres) GetAllArticles(ctx context.Context) ([]entity.Article, error) {
	query := fmt.Sprintf(`
		SELECT a.id, a.title, a.description, a.user_id, jsonb_agg(jsonb_build_object('id', c.id, 'name', c.name)) as categories 
		FROM %s a
		INNER JOIN %s ac ON a.id = ac.article_id
		INNER JOIN %s c ON c.id = ac.category_id
		GROUP BY a.user_id, a.id
	`, articlesTable, articlesCategoriesTable, categoriesTable)

	rows, err := p.Pool.Query(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var articles []entity.Article

	for rows.Next() {
		article := entity.Article{}
		var categoriesData []byte

		err := rows.Scan(
			&article.ID,
			&article.Title,
			&article.Description,
			&article.UserID,
			&categoriesData,
		)
		if err != nil {
			return nil, err
		}

		var categories []*entity.Category
		err = json.Unmarshal(categoriesData, &categories)
		if err != nil {
			return nil, err
		}

		article.Categories = make([]entity.Category, len(categories))
		for i, c := range categories {
			article.Categories[i] = *c
		}

		articles = append(articles, article)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return articles, nil
}

func (p *Postgres) GetArticlesByUserID(ctx context.Context, userID int64) ([]entity.Article, error) {
	query := fmt.Sprintf(`
		SELECT a.id, a.title, a.description, a.user_id, jsonb_agg(jsonb_build_object('id', c.id, 'name', c.name)) as categories 
		FROM %s a
		INNER JOIN %s ac ON a.id = ac.article_id
		INNER JOIN %s c ON c.id = ac.category_id
		WHERE a.user_id = $1
		GROUP BY a.user_id, a.id
	`, articlesTable, articlesCategoriesTable, categoriesTable)

	rows, err := p.Pool.Query(ctx, query, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var articles []entity.Article

	for rows.Next() {
		article := entity.Article{}
		var categoriesData []byte

		err := rows.Scan(
			&article.ID,
			&article.Title,
			&article.Description,
			&article.UserID,
			&categoriesData,
		)
		if err != nil {
			return nil, err
		}

		var categories []*entity.Category
		err = json.Unmarshal(categoriesData, &categories)
		if err != nil {
			return nil, err
		}

		article.Categories = make([]entity.Category, len(categories))
		for i, c := range categories {
			article.Categories[i] = *c
		}

		articles = append(articles, article)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return articles, nil
}

func (p *Postgres) GetUserIdByArticleId(ctx context.Context, articleId int64) (int64, error) {
	var userId int64

	query := fmt.Sprintf(
		`SELECT user_id from %s 
				WHERE id = $1`, articlesTable)

	row := p.Pool.QueryRow(ctx, query, articleId)

	err := row.Scan(&userId)
	if err != nil {
		return userId, err
	}

	return userId, nil
}
