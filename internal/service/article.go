package service

import (
	"api-blog/internal/entity"
	"context"
	"fmt"
)

func (m *Manager) CreateArticle(ctx context.Context, a *entity.Article) error {
	return m.Repository.CreateArticle(ctx, a)
}

func (m *Manager) UpdateArticle(ctx context.Context, a *entity.Article) error {
	return m.Repository.UpdateArticle(ctx, a)
}

func (m *Manager) DeleteArticle(ctx context.Context, id int64, userId int64) error {
	articleUserId, err := m.Repository.GetUserIdByArticleId(ctx, id)
	if err != nil {
		return err
	}

	if articleUserId != userId {
		return fmt.Errorf("permission denied")
	}

	return m.Repository.DeleteArticleByID(ctx, id)
}

func (m *Manager) GetArticleByID(ctx context.Context, id int64) (*entity.Article, error) {
	return m.Repository.GetArticleByID(ctx, id)
}

func (m *Manager) GetAllArticles(ctx context.Context) ([]*entity.Article, error) {
	return m.Repository.GetAllArticles(ctx)
}

func (m *Manager) GetArticlesByUserID(ctx context.Context, userID int64) ([]*entity.Article, error) {
	return m.Repository.GetArticlesByUserID(ctx, userID)
}
