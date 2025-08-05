package repo

import (
	"context"
	"database/sql"
	"devhive/internal/domain/models"
)

type DiscRepoInterface interface {
	CreateDisc(c context.Context, desc *models.Discussion) error
}

type DiscRepo struct {
	db *sql.DB
}

func NewDiscRepo(db *sql.DB) *DiscRepo {
	return &DiscRepo{db: db}
}

func (r *DiscRepo) CreateDisc(c context.Context, desc *models.Discussion) error {
	query := `INSERT INTO discussions (title, description, creator_id, created_at) VALUES ($1, $2, $3, $4) RETURNING id`

	return r.db.QueryRowContext(
		c,
		query,
		desc.Title,
		desc.Description,
		desc.CreatorID,
		desc.CreatedAt,
	).Scan(&desc.ID)
}
