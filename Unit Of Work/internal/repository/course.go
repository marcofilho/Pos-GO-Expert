package repository

import (
	"context"
	"database/sql"

	"github.com/marcofilho/Pos-GO-Expert/UnitOfWork/internal/db"
	"github.com/marcofilho/Pos-GO-Expert/UnitOfWork/internal/entity"
)

type CourseRepositoryInterface interface {
	Insert(ctx context.Context, course entity.Course) error
}

type CourseRepository struct {
	DB      *sql.DB
	Queries *db.Queries
}

func NewCourseRepository(dtb *sql.DB) *CourseRepository {
	return &CourseRepository{
		DB:      dtb,
		Queries: db.New(dtb),
	}
}

func (r *CourseRepository) Insert(ctx context.Context, course entity.Course) error {
	return r.Queries.CreateCourse(ctx, db.CreateCourseParams{
		Name:       course.Name,
		CategoryID: int32(course.CategoryID),
	})
}
