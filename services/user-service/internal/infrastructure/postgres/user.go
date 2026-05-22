package postgres

import (
	"context"
	"time"
	"users/internal/domain"
	usecases "users/internal/usecases/user"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgxpool"
)

type userModel struct {
	ID uuid.UUID

	Username string
	Email    string

	Password string

	CreatedAt time.Time
}

type UserRepository struct {
	db *pgxpool.Pool
}

func NewUserRepository(db *pgxpool.Pool) usecases.UserRepository {
	return &UserRepository{
		db: db,
	}
}

func (r *UserRepository) Create(ctx context.Context, user *domain.User) error {
	model := userEntityToModel(user)

	query := `
		INSERT INTO public.users (
			id,
			username,
			email,
			created_at
		)
		VALUES ($1, $2, $3, $4)
	`

	_, err := r.db.Exec(
		ctx,
		query,
		model.ID,
		model.Username,
		model.Email,
		model.CreatedAt,
	)

	if err != nil {
		return err
	}

	return nil
}

func (r *UserRepository) GetByID(ctx context.Context, id uuid.UUID) (*domain.User, error) {
	var model userModel

	query := `
		SELECT id, username, email, created_at
		FROM public.users
		WHERE id = $1
	`

	err := r.db.QueryRow(ctx, query, id).Scan(
		&model.ID,
		&model.Username,
		&model.Email,
		&model.CreatedAt,
	)

	if err != nil {
		return nil, err
	}

	return userModelToEntity(&model), nil
}

func (r *UserRepository) UpdateByID(ctx context.Context, user *domain.User) error {
	model := userEntityToModel(user)

	query := `
		UPDATE public.users
		SET username = $1, email = $2
		WHERE id = $3
	`

	_, err := r.db.Exec(
		ctx,
		query,
		model.Username,
		model.Email,
		model.ID,
	)

	if err != nil {
		return err
	}

	return nil
}

func (r *UserRepository) DeleteByID(ctx context.Context, id uuid.UUID) error {
	query := `
		DELETE FROM public.users
		WHERE id = $1
	`

	_, err := r.db.Exec(ctx, query, id)
	if err != nil {
		return err
	}

	return nil
}

func (r *UserRepository) List(ctx context.Context) ([]*domain.User, error) {
	query := `
		SELECT id, username, email, created_at
		FROM public.users
	`

	rows, err := r.db.Query(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []*domain.User
	for rows.Next() {
		var model userModel
		err := rows.Scan(
			&model.ID,
			&model.Username,
			&model.Email,
			&model.CreatedAt,
		)
		if err != nil {
			return nil, err
		}
		users = append(users, userModelToEntity(&model))
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}

	return users, nil
}

func userEntityToModel(entity *domain.User) *userModel {
	model := userModel{
		ID: entity.ID,

		Username: entity.Username,
		Email:    entity.Email,

		CreatedAt: entity.CreatedAt,
	}
	return &model
}

func userModelToEntity(model *userModel) *domain.User {
	entiry := domain.User{
		ID: model.ID,

		Username: model.Username,
		Email:    model.Email,
	}
	return &entiry
}
