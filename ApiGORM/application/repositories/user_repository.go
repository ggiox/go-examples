package repositories

import (
	"log"

	"ApiGORM/domain"
	"github.com/jinzhu/gorm"
)

// UserRepository represents a interface of User
type UserRepository interface {
	Insert(user *domain.User) (*domain.User, error)
}

// UserRepositoryDb represents a db conection
type UserRepositoryDb struct {
	Db *gorm.DB
}

// Insert fuction return err if no insert user in db
func (repo UserRepositoryDb) Insert(user *domain.User) (*domain.User, error) {

	err := user.Prepare()
	if err != nil {
		log.Fatalf("Error during the user validation: %v", err)
		return user, err
	}

	err = repo.Db.Create(user).Error
	if err != nil {
		log.Fatalf("Error to persist user: %v", err)
		return user, err
	}

	return user, nil

}
