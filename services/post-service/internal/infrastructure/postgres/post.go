package postgres

import (
	"context"
	"posts/internal/domain"
	usecases "posts/internal/usecases/post"
	"time"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgxpool"
)

type postModel struct {
	ID uuid.UUID

	Title    string
	Content  string
	AuthorID uuid.UUID

	CreatedAt time.Time
}

type PostRepository struct {
	db *pgxpool.Pool
}

func NewPostRepository(db *pgxpool.Pool) usecases.PostRepository {
	return &PostRepository{
		db: db,
	}
}

func (r *PostRepository) Create(ctx context.Context, post *domain.Post) error {
	model := postEntityToModel(post)

	query := `
		INSERT INTO public.posts (
			id,
			title,
			content,
			author_id,
			created_at
		)
		VALUES ($1, $2, $3, $4, $5)
	`

	_, err := r.db.Exec(
		ctx,
		query,
		model.ID,
		model.Title,
		model.Content,
		model.AuthorID,
		model.CreatedAt,
	)

	if err != nil {
		return err
	}

	return nil
}

func (r *PostRepository) GetByID(ctx context.Context, id uuid.UUID) (*domain.Post, error) {
	var model postModel

	query := `
		SELECT id, title, content, author_id, created_at
		FROM public.posts
		WHERE id = $1
	`

	err := r.db.QueryRow(ctx, query, id).Scan(
		&model.ID,
		&model.Title,
		&model.Content,
		&model.AuthorID,
		&model.CreatedAt,
	)

	if err != nil {
		return nil, err
	}

	return postModelToEntity(&model), nil
}

func (r *PostRepository) UpdateByID(ctx context.Context, post *domain.Post) error {
	model := postEntityToModel(post)

	query := `
		UPDATE public.posts
		SET title = $1, content = $2
		WHERE id = $3
	`

	_, err := r.db.Exec(
		ctx,
		query,
		model.Title,
		model.Content,
		model.ID,
	)

	if err != nil {
		return err
	}

	return nil
}

func (r *PostRepository) DeleteByID(ctx context.Context, id uuid.UUID) error {
	query := `
		DELETE FROM public.posts
		WHERE id = $1
	`

	_, err := r.db.Exec(ctx, query, id)
	if err != nil {
		return err
	}

	return nil
}

func (r *PostRepository) List(ctx context.Context) ([]*domain.Post, error) {
	query := `
		SELECT id, title, content, author_id, created_at	
		FROM public.posts
	`

	rows, err := r.db.Query(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var posts []*domain.Post
	for rows.Next() {
		var model postModel
		err := rows.Scan(
			&model.ID,
			&model.Title,
			&model.Content,
			&model.AuthorID,
			&model.CreatedAt,
		)
		if err != nil {
			return nil, err
		}
		posts = append(posts, postModelToEntity(&model))
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}

	return posts, nil
}

func postEntityToModel(entity *domain.Post) *postModel {
	model := postModel{
		ID: entity.ID,

		Title:     entity.Title,
		Content:   entity.Content,
		AuthorID:  entity.AuthorID,
		CreatedAt: entity.CreatedAt,
	}
	return &model
}

func postModelToEntity(model *postModel) *domain.Post {
	entity := domain.Post{
		ID: model.ID,

		Title:     model.Title,
		Content:   model.Content,
		AuthorID:  model.AuthorID,
		CreatedAt: model.CreatedAt,
	}
	return &entity
}
