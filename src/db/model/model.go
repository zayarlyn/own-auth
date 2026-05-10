// Package model
package model

type User struct {
	ID string `gorm:"type:uuid,primaryKey"`
}

func (User) TableName() string {
	return "user"
}
