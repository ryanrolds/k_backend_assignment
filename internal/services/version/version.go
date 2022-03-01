package version

import (
	"context"
	"database/sql"

	"kong-assignment.network/service_backend/internal/domain"
	"kong-assignment.network/service_backend/internal/persistence"
)

func GetVersionsForService(ctx context.Context, db *sql.DB, serviceID string) ([]domain.Version, error) {
	return persistence.GetVersionsForService(ctx, db, serviceID)
}
