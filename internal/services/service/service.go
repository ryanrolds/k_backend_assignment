package service

import (
	"context"
	"database/sql"
	"fmt"

	"kong-assignment.network/service_backend/internal/domain"
	"kong-assignment.network/service_backend/internal/persistence"
	"kong-assignment.network/service_backend/internal/services/version"
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
	service, err := persistence.GetService(ctx, db, id)
	if err != nil {
		return nil, fmt.Errorf("problem fetching service: %w", err)
	}

	if service == nil {
		return nil, fmt.Errorf("service not found")
	}

	versions, err := version.GetVersionsForService(ctx, db, id)
	if err != nil {
		return nil, fmt.Errorf("problem fetching versions: %w", err)
	}

	service.Versions = versions
	service.VersionCount = len(versions)

	return service, nil
}
