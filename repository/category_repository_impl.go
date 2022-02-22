package repository

import (
	"context"
	"database/sql"
	"errors"

	"github.com/sumitroajiprabowo/go-restfull-api/helper"
	"github.com/sumitroajiprabowo/go-restfull-api/model/entity"
)

type CategoryRepositoryImpl struct {
}

/*
Create a new instance of CategoryRepositoryImpl and return its pointer to the caller function (CategoryRepository) in order to be used in the service package.
*/
func NewCategoryRepository() CategoryRepository {
	return &CategoryRepositoryImpl{}
}

// CategoryRepositoryImplementation.Insert
func (c *CategoryRepositoryImpl) Create(ctx context.Context, tx *sql.Tx, category entity.Category) entity.Category {
	q := "insert into category (name) values (?)"
	result, err := tx.ExecContext(ctx, q, category.Name)
	helper.PanicIfError(err)

	id, err := result.LastInsertId()
	helper.PanicIfError(err)

	category.Id = id

	return category

}

// CategoryRepositoryImplementation.Update
func (c *CategoryRepositoryImpl) Update(ctx context.Context, tx *sql.Tx, category entity.Category) entity.Category {
	q := "update category set name = ? where id = ?"
	_, err := tx.ExecContext(ctx, q, category.Name, category.Id)
	helper.PanicIfError(err)

	return category
}

// CategoryRepositoryImplementation.Delete
func (c *CategoryRepositoryImpl) Delete(ctx context.Context, tx *sql.Tx, category entity.Category) {
	q := "delete from category where id = ?"
	_, err := tx.ExecContext(ctx, q, category.Id)
	helper.PanicIfError(err)
}

// CategoryRepositoryImplementation.FindById
func (c *CategoryRepositoryImpl) FindById(ctx context.Context, tx *sql.Tx, categoryId int64) (entity.Category, error) {
	q := "select id, name from category where id = ?"

	row, err := tx.QueryContext(ctx, q, categoryId)
	helper.PanicIfError(err)
	defer row.Close()

	category := entity.Category{}

	if row.Next() {
		err := row.Scan(&category.Id, &category.Name)
		helper.PanicIfError(err)
		return category, nil
	} else {
		return category, errors.New("Category not found")
	}
}

// CategoryRepositoryImplementation.FindAll
func (c *CategoryRepositoryImpl) FindAll(ctx context.Context, tx *sql.Tx) []entity.Category {
	q := "select id, name from category"

	rows, err := tx.QueryContext(ctx, q)
	helper.PanicIfError(err)
	defer rows.Close()

	var categories []entity.Category
	for rows.Next() {
		category := entity.Category{}
		err := rows.Scan(&category.Id, &category.Name)
		helper.PanicIfError(err)
		categories = append(categories, category)
	}

	return categories
}
