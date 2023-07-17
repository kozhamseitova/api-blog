package service

import (
	"api-blog/internal/entity"
	"context"
)

func (m *Manager) GetCategories(ctx context.Context) ([]entity.Category, error) {
	return m.Repository.GetCategories(ctx)
}
