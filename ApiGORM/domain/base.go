package domain

import (
	"log"
	"time"

	"github.com/google/uuid"
	"github.com/jinzhu/gorm"
)

// A Base represents the common fields for a data structure
type Base struct {
	ID        string    `json:"id" gorm:"type:uuid;primary_key"`
	CreatedAt time.Time `json:"create_at"`
	UpdateAt  time.Time `json:"update_at"`
	DeletedAt time.Time `json:"delete_at" sql:"index"`
}

// BeforeCreate function return error if not fill fields
func (base *Base) BeforeCreate(scope *gorm.Scope) error {

	err := scope.SetColumn("CreatedAt", time.Now())
	if err != nil {
		log.Fatalf("Error during CreateAt object creating")
		return err
	}

	err = scope.SetColumn("Id", uuid.NewString())
	if err != nil {
		log.Fatalf("Error during ID object creating")
		return err
	}

	return nil

}
