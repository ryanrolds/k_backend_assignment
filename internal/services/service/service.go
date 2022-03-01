package service

import (
	"context"
	"database/sql"

	"kong-assignment.network/service_backend/internal/domain"
	"kong-assignment.network/service_backend/internal/persistence"
)

func Page(ctx context.Context, db *sql.DB, offset int, limit int) ([]domain.Service, error) {
	return persistence.GetPageOfServices(ctx, db, offset, limit)
}

func PageSearched(ctx context.Context, db *sql.DB, search string, offset int, limit int) ([]domain.Service, error) {
	return persistence.GetPageOfSearchedServices(ctx, db, search, offset, limit)
}

func Count(ctx context.Context, db *sql.DB) (int, error) {
	return persistence.CountServices(ctx, db)
}

func Get(ctx context.Context, db *sql.DB, id string) (*domain.Service, error) {
	return persistence.GetService(ctx, db, id)
}
