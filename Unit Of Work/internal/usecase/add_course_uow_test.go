package usecase

import (
	"context"
	"database/sql"
	"testing"

	"github.com/marcofilho/Pos-GO-Expert/UnitOfWork/internal/db"
	"github.com/marcofilho/Pos-GO-Expert/UnitOfWork/internal/repository"
	"github.com/marcofilho/Pos-GO-Expert/UnitOfWork/pkg/uow"
	"github.com/stretchr/testify/assert"

	_ "github.com/go-sql-driver/mysql"
)

func TestAddCourseUnitOfWork(t *testing.T) {
	dbt, err := sql.Open("mysql", "root:root@tcp(localhost:3306)/goexpert")
	assert.NoError(t, err)

	dbt.Exec("DROP TABLE if exists `courses`;")
	dbt.Exec("DROP TABLE if exists `categories`;")

	dbt.Exec("CREATE TABLE IF NOT EXISTS `categories` (id int PRIMARY KEY AUTO_INCREMENT, name varchar(255) NOT NULL);")
	dbt.Exec("CREATE TABLE IF NOT EXISTS `courses` (id int PRIMARY KEY AUTO_INCREMENT, name varchar(255) NOT NULL, category_id INTEGER NOT NULL, FOREIGN KEY (category_id) REFERENCES categories(id));")

	ctx := context.Background()
	uow := uow.NewUnitOfWork(ctx, dbt)

	uow.Register("CategoryRepository", func(tx *sql.Tx) interface{} {
		repository := repository.NewCategoryRepository(dbt)
		repository.Queries = db.New(tx)
		return repository
	})

	uow.Register("CourseRepository", func(tx *sql.Tx) interface{} {
		repository := repository.NewCourseRepository(dbt)
		repository.Queries = db.New(tx)
		return repository
	})

	input := InputUseCase{
		CategoryName:     "Category 1",
		CourseName:       "Course 1",
		CourseCategoryID: 1,
	}

	useCase := NewAddCourseUseCaseUow(uow)
	err = useCase.Execute(ctx, input)
	assert.NoError(t, err)
}
