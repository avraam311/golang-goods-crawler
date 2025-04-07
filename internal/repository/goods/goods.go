package goodsrepo

import (
	"context"
	"errors"
	"fmt"

	"github.com/avraam311/golang-goods-crawler/internal/models/domain"
	"github.com/avraam311/golang-goods-crawler/internal/models/dto"
	"github.com/jackc/pgx"
	_ "github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type Repo struct {
	pool *pgxpool.Pool
}

func New(db *pgxpool.Pool) *Repo {
	return &Repo{
		pool: db,
	}
}

func (r *Repo) Create(ctx context.Context, good dto.Goods) error {
	var id int
	query := `
		INSERT INTO goods (name)
		VALUES ($1)
		RETURNING id;
	`
	if err := r.pool.QueryRow(ctx, query, good.Name).Scan(&id); err != nil {
		return fmt.Errorf("repository.Goods.Create: %w", err)
	}
	return nil
}

func (r *Repo) GetAll(ctx context.Context) ([]domain.Goods, error) {
	var towns []domain.Goods
	query := `
		SELECT c.id, c.name  
		FROM goods c;
	`
	rows, err := r.pool.Query(ctx, query)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, nil
		}
		return nil, fmt.Errorf("repo.Goods.GetAll: %w", err)
	}
	for rows.Next() {
		town := domain.Goods{}
		err := rows.Scan(
			&town.ID,
			&town.Name,
		)
		if err != nil {
			return nil, fmt.Errorf("repository.Goods.GetAll: %w", err)
		}
		towns = append(towns, town)
	}
	return towns, nil
}
