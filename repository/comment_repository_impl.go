package repository

import (
	"context"
	"database/sql"
	"errors"
	"golang-database/entity"
	"strconv"
)

// ini adalah implementasi dari interface repository
type commentRepositoryImpl struct {
	DB *sql.DB
}

// buat method untuk membuat implementasi dari CommentRepositoryImpl
func NewCommentRepository(db *sql.DB) CommentRepository {
	return &commentRepositoryImpl{DB: db}
}

// function untuk interface
func (repository *commentRepositoryImpl) Insert(ctx context.Context, comment entity.Comment) (entity.Comment, error) {
	script := "INSERT INTO comments(email, comment) VALUES (?, ?)"
	result, err := repository.DB.ExecContext(ctx, script, comment.Email, comment.Commnent)
	if err != nil {
		// balikan comment tapi error
		return comment, err
	}
	// ambil id terakhir
	id, err := result.LastInsertId()
	if err != nil {
		return comment, err
	}
	comment.Id = int32(id)
	return comment, nil
}

func (repository *commentRepositoryImpl) FindById(ctx context.Context, id int32) (entity.Comment, error) {
	script := "SELECT id, email, comment FROM comments WHERE id = ? LIMIT 1"
	rows, err := repository.DB.QueryContext(ctx, script, id)
	comment := entity.Comment{}
	if err != nil {
		return comment, err
	}
	defer rows.Close()
	if rows.Next() {
		// ada
		rows.Scan(&comment.Id, &comment.Email, &comment.Commnent)
		return comment, nil
	} else {
		// tidak ada
		return comment, errors.New("Id " + strconv.Itoa(int(id)) + " Not Found")
	}
}

func (repository *commentRepositoryImpl) FindAll(ctx context.Context) ([]entity.Comment, error) {

	script := "SELECT id, email, comment FROM comments"
	rows, err := repository.DB.QueryContext(ctx, script)

	if err != nil {
		return nil, err
	}

	defer rows.Close()
	var comments []entity.Comment
	for rows.Next() {
		comment := entity.Comment{}
		rows.Scan(&comment.Id, &comment.Email, &comment.Commnent)
		comments = append(comments, comment)
	}
	return comments, nil
}
