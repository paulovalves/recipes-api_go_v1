package ingredient

import (
	"context"
	"errors"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/paulovalves/recipes-api/internal/domain"
)

var (
	ErrDuplicateName = errors.New("[Ingrediente]: já existe um ingrediente com este nome")
	ErrNotFound      = errors.New("[Ingrediente]: ingrediente não encontrado")
)

type Repository interface {
	Save(ctx context.Context, ingredient domain.Ingredient) error
	GetById(ctx context.Context, id int) (domain.Ingredient, error)
	GetAll(ctx context.Context) ([]domain.Ingredient, error)
	SearchByName(ctx context.Context, name string) (domain.Ingredient, error)
	Update(ctx context.Context, ingredient domain.Ingredient) error
	Delete(ctx context.Context, id int) error
}

type repository struct {
	db *pgxpool.Pool
}
