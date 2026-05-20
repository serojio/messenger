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

	PasswordHash string

	CreatedAt time.Time
	UpdatedAt time.Time
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
		INSERT INTO users (
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
