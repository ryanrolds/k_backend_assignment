package persistence

import (
	"context"
	"database/sql"

	"kong-assignment.network/service_backend/internal/domain"
)

func GetVersionsForService(ctx context.Context, db *sql.DB, serviceId string) ([]domain.Version, error) {
	versions := []domain.Version{}

	rows, err := db.Query("SELECT id, service_id, created_at, updated_at FROM version WHERE service_id = $1", serviceId)
	if err != nil {
		if err == sql.ErrNoRows {
			return versions, nil
		}

		return versions, err
	}
	defer rows.Close()

	for rows.Next() {
		version := domain.Version{}

		err := rows.Scan(&version.Id, &version.ServiceId, &version.CreatedAt, &version.UpdatedAt)
		if err != nil {
			return versions, err
		}

		versions = append(versions, version)
	}

	return versions, nil
}
