package generic

import (
	"autumn/pkg/autumncore/connection"
	"os"
)

type Model[T any] interface {
	Get() ([]T, error)
	Insert(data T) error
	Update(T) error
	Delete(interface{}) error
}

type ModelImpl[T Model[T]] struct {
}

func (m *ModelImpl[T]) Get() ([]T, error) {
	db, err := connection.NewDatabase().GetConnection(os.Getenv("DATABASE_DSN"))
	if err != nil {
		return nil, err
	}

	defer db.Close()

	var data []T

	tx := db.GormQuery().Find(&data)
	if tx.Error != nil {
		return nil, tx.Error
	}

	return data, nil
}

func (m *ModelImpl[T]) Insert(data T) error {
	db, err := connection.NewDatabase().GetConnection(os.Getenv("DATABASE_DSN"))
	if err != nil {
		return err
	}

	defer db.Close()

	tx := db.GormQuery().Create(data)
	if tx.Error != nil {
		return tx.Error
	}

	return nil
}

func (m *ModelImpl[T]) Update(dataUpdate T) error {
	db, err := connection.NewDatabase().GetConnection(os.Getenv("DATABASE_DSN"))
	if err != nil {
		return err
	}

	defer db.Close()

	tx := db.GormQuery().Save(dataUpdate)
	if tx.Error != nil {
		return tx.Error
	}

	return nil
}

func (m *ModelImpl[T]) Delete(pkid interface{}) error {
	db, err := connection.NewDatabase().GetConnection(os.Getenv("DATABASE_DSN"))
	if err != nil {
		return err
	}

	defer db.Close()

	var data T

	tx := db.GormQuery().First(&data, pkid)
	if tx.Error != nil {
		return tx.Error
	}

	tx = db.GormQuery().Delete(&data, pkid)
	if tx.Error != nil {
		return tx.Error
	}

	return nil
}
