// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.24.0

package databse

import (
	"time"

	"github.com/google/uuid"
)

type NewUser struct {
	ID        uuid.UUID
	CreatedAt time.Time
	UpdatedAt time.Time
	Name      string
}
