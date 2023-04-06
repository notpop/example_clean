package repository

import (
	// "database/sql"
	"example_clean/entity"
)

type UserRepository interface {
	FindByID(id string) (*entity.User, error)
}

type userRepositoryImpl struct {
	// db *sql.DB
}

// func NewUserRepository(db *sql.DB) UserRepository {
// 	return &userRepositoryImpl{db: db}
// }

func NewUserRepository() UserRepository {
	return &userRepositoryImpl{}
}

func (ur *userRepositoryImpl) FindByID(id string) (*entity.User, error) {
	// query := "SELECT id, name, email FROM users WHERE id = ?"
	// row := ur.db.QueryRow(query, id)

	// var user entity.User
	// err := row.Scan(&user.ID, &user.Name, &user.Email)
	// if err != nil {
	// 	return nil, err
	// }

	// return &user, nil

	return &entity.User{
		ID:   "1",
		Name: "John Doe",
	}, nil
}
