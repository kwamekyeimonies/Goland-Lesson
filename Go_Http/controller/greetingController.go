package controller

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/google/uuid"
	"github.com/kwamekyeimonies/Go-Http/database"
	"github.com/kwamekyeimonies/Go-Http/models"
)

func CreateUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	db := database.Database_Connection()
	defer db.Close()

	user := &models.User{
		ID: uuid.New().String(),
	}

	_ = json.NewDecoder(r.Body).Decode(&user)

	_, err := db.Model(user).Insert()
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	json.NewEncoder(w).Encode(user)

}

func CreateProduct(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	db := database.Database_Connection()
	defer db.Close()

	product := &models.Product{
		ID: uuid.New().String(),
	}

	_ = json.NewDecoder(r.Body).Decode(&product)

	_, err := db.Model(product).Insert()
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	json.NewEncoder(w).Encode(product)

}
