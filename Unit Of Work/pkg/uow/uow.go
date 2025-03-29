package uow

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/pkg/errors"
)

type RepositoryFactory func(tx *sql.Tx) interface{}

type UnitOfWorkInterface interface {
	Register(name string, fc RepositoryFactory)
	GetRepository(ctx context.Context, name string) (interface{}, error)
	Do(ctx context.Context, fn func(uow *UnitOfWork) error) error
	CommitOrRollback() error
	Rollback() error
	UnRegister(name string)
}

type UnitOfWork struct {
	Db           *sql.DB
	Tx           *sql.Tx
	Repositories map[string]RepositoryFactory
}

func NewUnitOfWork(ctx context.Context, db *sql.DB) *UnitOfWork {
	return &UnitOfWork{
		Db:           db,
		Repositories: make(map[string]RepositoryFactory),
	}
}

func (uow *UnitOfWork) Register(name string, fc RepositoryFactory) {
	uow.Repositories[name] = fc
}

func (uow *UnitOfWork) UnRegister(name string) {
	delete(uow.Repositories, name)
}

func (uow *UnitOfWork) GetRepository(ctx context.Context, name string) (interface{}, error) {
	if uow.Tx == nil {
		tx, err := uow.Db.BeginTx(ctx, nil)
		if err != nil {
			return nil, err
		}
		uow.Tx = tx
	}

	repo := uow.Repositories[name](uow.Tx)
	return repo, nil
}

func (uow *UnitOfWork) Do(ctx context.Context, fn func(uow *UnitOfWork) error) error {
	if uow.Tx != nil {
		return fmt.Errorf("transaction already started")
	}

	tx, err := uow.Db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}

	uow.Tx = tx
	err = fn(uow)
	if err != nil {
		errRb := uow.Rollback()
		if errRb != nil {
			return errors.New(fmt.Sprintf("original error: %s, rollback error: %s", err.Error(), errRb.Error()))
		}
		return err
	}
	return uow.CommitOrRollback()
}

func (u *UnitOfWork) Rollback() error {
	if u.Tx == nil {
		return errors.New("no transaction to rollback")
	}

	err := u.Tx.Rollback()
	if err != nil {
		return err
	}

	u.Tx = nil
	return nil
}

func (u *UnitOfWork) CommitOrRollback() error {
	err := u.Tx.Commit()
	if err != nil {
		errRb := u.Rollback()
		if errRb != nil {
			return errors.New(fmt.Sprintf("original error: %s, rollback error: %s", err.Error(), errRb.Error()))
		}
		return err
	}

	u.Tx = nil
	return nil
}
