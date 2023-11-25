package main

import (
	"gihub/com/nurmuhammaddeveloper/rssag/internal/databse"
	"time"

	"github.com/google/uuid"
)

type User struct {
	Id        uuid.UUID `json:"id"`
	CreatedAt time.Time `json:"creeated_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Name      string    `json:"name"`
}

func databaseUserToUser(db databse.NewUser  ) User{
	return User{
		Id: db.ID,
		CreatedAt: db.CreatedAt,
		UpdatedAt: db.UpdatedAt,
		Name: db.Name,
	}
}
