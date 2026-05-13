// Package model
package model

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/datatypes"
	"gorm.io/gorm"
)

type User struct {
	// Identity
	ID     uuid.UUID `gorm:"type:uuid;primaryKey;default:gen_random_uuid()" json:"id"`
	PoolID uuid.UUID `gorm:"type:uuid;not null"                             json:"poolId"`

	// Credentials
	Email        string `gorm:"type:text;not null" json:"email"`
	Username     string `gorm:"type:text;not null" json:"username"`
	PasswordHash string `gorm:"type:text;not null" json:"-"` // never expose

	// Status
	Status     string         `gorm:"type:text;not null;default:'UNCONFIRMED'" json:"status"`
	Attributes datatypes.JSON `gorm:"type:jsonb;not null;default:'{}'"         json:"attributes"`

	// Timestamps (nullable)
	LastLoginAt     *time.Time `gorm:"default:null" json:"lastLoginAt"`
	EmailVerifiedAt *time.Time `gorm:"default:null" json:"emailVerifiedAt"`

	// Timestamps
	CreatedAt time.Time      `gorm:"not null;autoCreateTime" json:"createdAt"`
	UpdatedAt time.Time      `gorm:"not null;autoUpdateTime" json:"updatedAt"`
	DeletedAt gorm.DeletedAt `gorm:"index"                   json:"deletedAt"`

	// Associations
	Pool UserPool `gorm:"foreignKey:PoolID;references:ID" json:"pool"`
}

func (User) TableName() string {
	return "user"
}
