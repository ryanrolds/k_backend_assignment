package persistence

import (
	"context"
	"database/sql"

	"kong-assignment.network/service_backend/internal/domain"
)

const servicesPage = `SELECT s.id, s.name, s.description, count(v.id) as version_count, s.created_at, s.updated_at ` +
	`FROM service s ` +
	`LEFT JOIN version as v ON s.id = v.service_id ` +
	`GROUP BY s.id ` +
	`ORDER BY s.updated_at ` +
	`OFFSET $1 ` +
	`LIMIT $2`

func GetPageOfServices(ctx context.Context, db *sql.DB, offset int, limit int) ([]domain.Service, error) {
	return queryServices(ctx, db, servicesPage, offset, limit)
}

func CountServices(ctx context.Context, db *sql.DB) (int, error) {
	var count int

	err := db.QueryRow("SELECT COUNT(*) FROM service").Scan(&count)
	if err != nil {
		return 0, err
	}

	return count, nil
}

const searchedServicesPage = `SELECT s.id, s.name, s.description, count(v.id) as version_count, s.created_at, s.updated_at ` +
	`FROM service s ` +
	`LEFT JOIN version as v ON s.id = v.service_id ` +
	`WHERE tsvector_name_description @@ to_tsquery($1) ` +
	`GROUP BY s.id ` +
	`ORDER BY ts_rank_cd(tsvector_name_description, to_tsquery($1)) DESC ` +
	`OFFSET $2 ` +
	`LIMIT $3`

func GetPageOfSearchedServices(ctx context.Context, db *sql.DB, search string, offset int, limit int) ([]domain.Service, error) {
	return queryServices(ctx, db, searchedServicesPage, search, offset, limit)
}

func CountSearchedServices(ctx context.Context, db *sql.DB, search string) (int, error) {
	var count int

	query := "SELECT COUNT(*) FROM service, to_tsquery($1) query WHERE tsvector_name_description @@ query"
	err := db.QueryRow(query, search).Scan(&count)
	if err != nil {
		return 0, err
	}

	return count, nil
}

func GetService(ctx context.Context, db *sql.DB, id string) (*domain.Service, error) {
	service := &domain.Service{}

	err := db.QueryRow("SELECT id, name, description, created_at, updated_at FROM service WHERE id = $1", id).
		Scan(&service.Id, &service.Name, &service.Description, &service.CreatedAt, &service.UpdatedAt)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}

		return nil, err
	}

	return service, nil
}

func queryServices(ctx context.Context, db *sql.DB, query string, args ...interface{}) ([]domain.Service, error) {
	services := []domain.Service{}

	rows, err := db.Query(query, args...)
	if err != nil {
		if err == sql.ErrNoRows {
			return services, nil
		}

		return services, err
	}
	defer rows.Close()

	for rows.Next() {
		service := domain.Service{}
		err := rows.Scan(&service.Id, &service.Name, &service.Description, &service.VersionCount, &service.CreatedAt, &service.UpdatedAt)
		if err != nil {
			return services, err
		}

		services = append(services, service)
	}

	err = rows.Err()
	if err != nil {
		return services, err
	}

	return services, nil
}
