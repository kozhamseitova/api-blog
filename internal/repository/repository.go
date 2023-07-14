package repository

import (
	"api-blog/internal/entity"
	"context"
)

type Repository interface {
	CreateUser(ctx context.Context, u *entity.User) error
	//Login(ctx context.Context, username, password string) (string, error)
	//UpdateUser(ctx context.Context, u *entity.User) error
	//DeleteUser(ctx context.Context, id int64) error
	//
	//CreateArticle(ctx context.Context, a *entity.Article) error
	//UpdateArticle(ctx context.Context, a *entity.Article) error
	//DeleteArticleByID(ctx context.Context, id int64) error
	//GetArticleByID(ctx context.Context, id int64) (*entity.Article, error)
	//GetAllArticles(ctx context.Context) ([]entity.Article, error)
	//GetArticlesByUserID(ctx context.Context, userID int64) ([]entity.Article, error)
	//
	//GetCategories(ctx context.Context) ([]entity.Category, error)
	//GetCategoryByID(ctx context.Context) (entity.Category, error)
}
