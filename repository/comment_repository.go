package repository

import (
	"context"
	"golang-database/entity"
)

// kita buat interface
type CommentRepository interface {
	// buat function-function yang digunakan untuk komunikasi dengan database

	// parameternya adalah context dan comment entity, balikannya adalah Comment dan error
	Insert(ctx context.Context, comment entity.Comment) (entity.Comment, error)

	FindById(ctx context.Context, id int32) (entity.Comment, error)

	FindAll(ctx context.Context) ([]entity.Comment, error)
}
