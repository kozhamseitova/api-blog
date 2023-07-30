package service

import (
	"context"
	"github.com/kozhamseitova/api-blog/internal/entity"
)

func (m *Manager) GetCategories(ctx context.Context) ([]*entity.Category, error) {
	return m.Repository.GetCategories(ctx)
}
