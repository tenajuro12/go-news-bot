package storage

import (
	"context"
	"go-news-bot/internal/model"
	"time"

	"github.com/jmoiron/sqlx"
	"github.com/samber/lo"
)

type SourcePostgresStorage struct {
	db *sqlx.DB
}

func NewSourceStorage(db *sqlx.DB) *SourcePostgresStorage {
	return &SourcePostgresStorage{db: db}
}

func (s *SourcePostgresStorage) Sources(ctx context.Context) ([]model.Source, error) {
	var sources []dbSource
	if err := s.db.SelectContext(ctx, &sources, `SELECT *FROM sources`); err != nil {
		return nil, err
	}

	return lo.Map(sources, func(source dbSource, _ int) model.Source {
		return model.Source{
			ID:        source.ID,
			Name:      source.Name,
			FeedUrl:   source.FeedURL,
			CreatedAt: source.CreatedAt,
		}
	}), nil

}

func (s *SourcePostgresStorage) SourceByID(ctx context.Context, id int64) (*model.Source, error) {
	var sources dbSource
	if err := s.db.GetContext(ctx, &sources, `SELECT *FROM sources WHERE id=$1`, id); err != nil {
		return nil, err
	}

	return &model.Source{
		ID:        sources.ID,
		Name:      sources.Name,
		FeedUrl:   sources.FeedURL,
		CreatedAt: sources.CreatedAt,
	}, nil
}

func (s *SourcePostgresStorage) Add(ctx context.Context, source model.Source) (int64, error) {
	var id int64
	err := s.db.QueryRowxContext(ctx, `INSERT INTO sources (name, feed_url, priority) 
VALUES ($1, $2, $3) RETURNING id;`, source.Name, source.FeedUrl,
	).Scan(&id)
	if err != nil {
		return 0, err
	}

	return id, nil
}

func (s *SourcePostgresStorage) Delete(ctx context.Context, id int64) error {
	_, err := s.db.ExecContext(ctx, `DELETE *FROM sources WHERE id=$1`, id)
	return err
}

type dbSource struct {
	ID        int64     `db:"id"`
	Name      string    `db:"name"`
	FeedURL   string    `db:"feed_url"`
	CreatedAt time.Time `db:"created_at"`
}
