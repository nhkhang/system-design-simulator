package models

import (
	"time"
)

type BaseModel struct {
	CreatedAt time.Time `json:"created_at"`
}
