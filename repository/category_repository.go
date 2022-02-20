package repository

import (
	"context"
	"database/sql"

	"github.com/sumitroajiprabowo/go-restfull-api/model/entity"
)

type CategoryRepository interface {
	Insert(ctx context.Context, tx *sql.Tx, category entity.Category) entity.Category
	Update(ctx context.Context, tx *sql.Tx, category entity.Category) entity.Category
	Delete(ctx context.Context, tx *sql.Tx, category entity.Category)
	FindById(ctx context.Context, tx *sql.Tx, categoryId int64) (entity.Category, error)
	FindAll(ctx context.Context, tx *sql.Tx) []entity.Category
}
