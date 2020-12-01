package base

import (
	"time"

	uuid "github.com/satori/go.uuid"
)

// Base struct for Structs
type Base struct {
	ID        uuid.UUID `gorm:"primary_key;type:uuid;"`
	CreatedAt time.Time
	UpdatedAt time.Time
}
