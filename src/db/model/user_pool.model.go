// Package model
package model

import (
	"time"

	"github.com/google/uuid"
	"github.com/lib/pq"
	"gorm.io/datatypes"
	"gorm.io/gorm"
)

type PasswordPolicy struct {
	MinLength        int  `json:"min_length"`
	RequireUppercase bool `json:"require_uppercase"`
	RequireLowercase bool `json:"require_lowercase"`
	RequireNumbers   bool `json:"require_numbers"`
	RequireSymbols   bool `json:"require_symbols"`
}

type UserPool struct {
	// Identity
	ID   uuid.UUID `gorm:"type:uuid;primaryKey;default:gen_random_uuid()" json:"id"`
	Name string    `gorm:"type:text;not null"                              json:"name"`

	// Token lifetimes
	TokenExpirySeconds int `gorm:"not null;default:900" json:"tokenExpirySeconds"`
	RefreshExpiryDays  int `gorm:"not null;default:30"  json:"refreshExpiryDays"`

	// Password policy (stored as JSONB)
	PasswordPolicy datatypes.JSONType[PasswordPolicy] `gorm:"type:jsonb;not null;default:'{}'" json:"passwordPolicy"`

	// Arrays
	AllowedFlows pq.StringArray `gorm:"type:text[];not null;default:'{\"email:password\"}'" json:"allowedFlows"`
	CallbackURLs pq.StringArray `gorm:"type:text[];not null;default:'{}'"                  json:"callbackUrls"`

	// Timestamps
	CreatedAt time.Time      `gorm:"not null;autoCreateTime" json:"createdAt"`
	UpdatedAt time.Time      `gorm:"not null;autoUpdateTime" json:"updatedAt"`
	DeletedAt gorm.DeletedAt `gorm:"index"                   json:"deletedAt"`
}

func (UserPool) TableName() string {
	return "user_pool"
}
