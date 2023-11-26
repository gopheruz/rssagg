package main

import (
	"encoding/json"
	"gihub/com/nurmuhammaddeveloper/rssag/internal/databse"
	"net/http"
	"time"

	"github.com/google/uuid"
)

func responseWithJsonHandler(w http.ResponseWriter, r *http.Request) {
	responseWithJson(w, 200, responsestruct{
		Message: "hammasi chotki ishlayapdi",
	})
}

func (apiCfg *apiConfig) handlerCreateuser(w http.ResponseWriter, r *http.Request) {
	type parametrs struct {
		Name string `json:"name"`
	}
	params := parametrs{}
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&params)
	if err != nil {
		responseWithJson(w, 400, responsestruct{
			Message: "Error parsing json",
		})
		return
	}
	user, err := apiCfg.Db.CreateUser(r.Context(), databse.CreateUserParams{
		ID:        uuid.New(),
		Name:      params.Name,
		CreatedAt: time.Now().UTC().Add(5),
		UpdatedAt: time.Now().UTC().Add(5),
	})
	if err != nil {
		responseWithJson(w, 400, responsestruct{
			Message: "Couldnt catch user" + err.Error(),
		})
		return
	}
	responseWithJson(w, 201, databaseUserToUser(user))
}
