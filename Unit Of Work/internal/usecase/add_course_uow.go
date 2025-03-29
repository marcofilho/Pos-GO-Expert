package usecase

import (
	"context"

	"github.com/marcofilho/Pos-GO-Expert/UnitOfWork/internal/entity"
	"github.com/marcofilho/Pos-GO-Expert/UnitOfWork/internal/repository"
	"github.com/marcofilho/Pos-GO-Expert/UnitOfWork/pkg/uow"
)

type InputUseCaseUow struct {
	CategoryName     string
	CourseName       string
	CourseCategoryID int
}

type AddCourseUseCaseUow struct {
	Uow uow.UnitOfWorkInterface
}

func NewAddCourseUseCaseUow(uow uow.UnitOfWorkInterface) *AddCourseUseCaseUow {
	return &AddCourseUseCaseUow{
		Uow: uow,
	}
}

func (a *AddCourseUseCaseUow) Execute(ctx context.Context, input InputUseCase) error {
	return a.Uow.Do(ctx, func(uow *uow.UnitOfWork) error {
		category := entity.Category{
			Name: input.CategoryName,
		}
		repoCategory := a.getCategoryRepository(ctx)
		err := repoCategory.Insert(ctx, category)
		if err != nil {
			return err
		}

		course := entity.Course{
			Name:       input.CourseName,
			CategoryID: input.CourseCategoryID,
		}

		repoCourse := a.getCourseRepository(ctx)
		err = repoCourse.Insert(ctx, course)
		if err != nil {
			return err
		}
		return nil
	})
}

func (a *AddCourseUseCaseUow) getCategoryRepository(ctx context.Context) repository.CategoryRepositoryInterface {
	repo, err := a.Uow.GetRepository(ctx, "CategoryRepository")
	if err != nil {
		panic(err)
	}
	return repo.(repository.CategoryRepositoryInterface)
}

func (a *AddCourseUseCaseUow) getCourseRepository(ctx context.Context) repository.CourseRepositoryInterface {
	repo, err := a.Uow.GetRepository(ctx, "CourseRepository")
	if err != nil {
		panic(err)
	}
	return repo.(repository.CourseRepositoryInterface)
}
