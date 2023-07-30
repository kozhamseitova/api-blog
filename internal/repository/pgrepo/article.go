package pgrepo

import (
	"context"
	"fmt"
	"github.com/georgysavva/scany/pgxscan"
	"github.com/kozhamseitova/api-blog/internal/entity"
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
		return fmt.Errorf("insert article query err: %w", err)
	}

	querySecond := fmt.Sprintf(`
		INSERT INTO %s (article_id, category_id)
		VALUES ($1, $2)`, articlesCategoriesTable)

	for _, category := range a.Categories {
		_, err = tx.Exec(ctx, querySecond, articleID, category.ID)
		if err != nil {
			_ = tx.Rollback(ctx) // Rollback the transaction on error
			return fmt.Errorf("insert article_categories query err: %w", err)
		}
	}

	err = tx.Commit(ctx)
	if err != nil {
		return fmt.Errorf("commit err: %w", err)
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
		return fmt.Errorf("query first err: %w", err)
	}

	deleteCategoriesQuery := fmt.Sprintf(
		`DELETE FROM %s
		WHERE article_id = $1`, articlesCategoriesTable)

	_, err = tx.Exec(ctx, deleteCategoriesQuery, a.ID)
	if err != nil {
		return fmt.Errorf("delete categories query err: %w", err)
	}

	insertCategoriesQuery := fmt.Sprintf(
		`INSERT INTO %s (article_id, category_id)
		VALUES ($1, $2)`, articlesCategoriesTable)

	for _, category := range a.Categories {
		_, err = tx.Exec(ctx, insertCategoriesQuery, a.ID, category.ID)
		return fmt.Errorf("insert categories query err: %w", err)
	}

	err = tx.Commit(ctx)
	if err != nil {
		return fmt.Errorf("commit err: %w", err)
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
		return fmt.Errorf("delete articles_categories query err: %w", err)
	}

	deleteArticleQuery := fmt.Sprintf(
		`DELETE FROM %s
						WHERE id = $1`, articlesTable)

	_, err = tx.Exec(ctx, deleteArticleQuery, id)
	if err != nil {
		return fmt.Errorf("delete articles query err: %w", err)
	}

	err = tx.Commit(ctx)
	if err != nil {
		return fmt.Errorf("commit err: %w", err)
	}

	return nil
}

func (p *Postgres) GetArticleByID(ctx context.Context, id int64) (*entity.Article, error) {
	article := new(entity.Article)

	query := fmt.Sprintf(
		`SELECT a.id, a.title, a.description, a.user_id, jsonb_agg(jsonb_build_object('id', c.id, 'name', c.name)) as categories from %s a
				inner join %s ac on a.id = ac.article_id
				inner join %s c on c.id = ac.category_id
				where a.id = $1
				group by a.user_id, a.id`, articlesTable, articlesCategoriesTable, categoriesTable)

	err := pgxscan.Get(ctx, p.Pool, article, query, id)

	if err != nil {
		return article, err
	}

	return article, nil
}

func (p *Postgres) GetAllArticles(ctx context.Context) ([]*entity.Article, error) {
	query := fmt.Sprintf(`
		SELECT a.id, a.title, a.description, a.user_id, jsonb_agg(jsonb_build_object('id', c.id, 'name', c.name)) as categories 
		FROM %s a
		INNER JOIN %s ac ON a.id = ac.article_id
		INNER JOIN %s c ON c.id = ac.category_id
		GROUP BY a.user_id, a.id
	`, articlesTable, articlesCategoriesTable, categoriesTable)

	var articles []*entity.Article

	err := pgxscan.Select(ctx, p.Pool, &articles, query)

	if err != nil {
		return nil, err
	}

	return articles, nil
}

func (p *Postgres) GetArticlesByUserID(ctx context.Context, userID int64) ([]*entity.Article, error) {
	var articles []*entity.Article

	query := fmt.Sprintf(`
		SELECT a.id, a.title, a.description, a.user_id, jsonb_agg(jsonb_build_object('id', c.id, 'name', c.name)) as categories 
		FROM %s a
		INNER JOIN %s ac ON a.id = ac.article_id
		INNER JOIN %s c ON c.id = ac.category_id
		WHERE a.user_id = $1
		GROUP BY a.user_id, a.id
	`, articlesTable, articlesCategoriesTable, categoriesTable)

	err := pgxscan.Select(ctx, p.Pool, &articles, query, userID)

	if err != nil {
		return nil, err
	}

	return articles, nil
}

func (p *Postgres) GetUserIdByArticleId(ctx context.Context, articleId int64) (int64, error) {
	var userId int64

	query := fmt.Sprintf(
		`SELECT user_id from %s 
				WHERE id = $1`, articlesTable)

	err := pgxscan.Get(ctx, p.Pool, &userId, query, articleId)

	if err != nil {
		return userId, fmt.Errorf("get user id by article id errror: %w", err)
	}

	return userId, nil
}
