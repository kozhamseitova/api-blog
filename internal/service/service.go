package service

import (
	"context"
	"github.com/kozhamseitova/api-blog/internal/entity"
)

type Service interface {
	CreateUser(ctx context.Context, u *entity.User) error
	Login(ctx context.Context, username, password string) (string, error)
	UpdateUser(ctx context.Context, u *entity.User) error
	DeleteUser(ctx context.Context, id int64) error
	VerifyToken(token string) (int64, error)

	CreateArticle(ctx context.Context, a *entity.Article) error
	UpdateArticle(ctx context.Context, a *entity.Article) error
	DeleteArticle(ctx context.Context, id int64, userId int64) error
	GetArticleByID(ctx context.Context, id int64) (*entity.Article, error)
	GetAllArticles(ctx context.Context) ([]*entity.Article, error)
	GetArticlesByUserID(ctx context.Context, userID int64) ([]*entity.Article, error)

	GetCategories(ctx context.Context) ([]*entity.Category, error)
}
