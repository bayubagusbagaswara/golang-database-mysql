package repository

import (
	"context"
	"fmt"
	golang_database "golang-database"
	"golang-database/entity"
	"testing"

	_ "github.com/go-sql-driver/mysql"
)

func TestCommentInsert(t *testing.T) {

	// NewCommentRepository adalah method implementasi dari CommentRepositoryImpl
	// seperti constructor
	commentRepository := NewCommentRepository(golang_database.GetConnection())

	ctx := context.Background()
	comment := entity.Comment{
		Email:    "repository@test.com",
		Commnent: "Test Repository",
	}

	result, err := commentRepository.Insert(ctx, comment)

	if err != nil {
		panic(err)
	}
	fmt.Println(result)
}

func TestFindById(t *testing.T) {
	commentRepository := NewCommentRepository(golang_database.GetConnection())

	id := 13
	ctx := context.Background()
	comment, err := commentRepository.FindById(ctx, int32(id))

	if err != nil {
		panic(err)
	}
	fmt.Println(comment)
}

func TestFindAll(t *testing.T) {
	commentRepository := NewCommentRepository(golang_database.GetConnection())

	ctx := context.Background()
	comments, err := commentRepository.FindAll(ctx)

	if err != nil {
		panic(err)
	}
	for _, comment := range comments {
		fmt.Println(comment)
	}
}
