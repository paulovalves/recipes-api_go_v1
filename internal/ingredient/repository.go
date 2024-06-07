package ingredient

import (
	"context"
	"errors"

	"github.com/paulovalves/recipes-api/internal/domain"
	"gorm.io/gorm"
)

var (
	ErrDuplicateName = errors.New("[Ingrediente]: já existe um ingrediente com este nome")
	ErrNotFound      = errors.New("[Ingrediente]: ingrediente não encontrado")
	ErrFound         = errors.New("[Ingrediente]: ingrediente encontrado")
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
	db *gorm.DB
}

func NewRepository(db *gorm.DB) Repository {
	return &repository{db: db}
}

func (r *repository) Save(ctx context.Context, ingredient domain.Ingredient) error {
	ingredientFound, err := r.SearchByName(ctx, ingredient.Name)
	if err != nil && err != ErrNotFound {
		return err
	}
}
