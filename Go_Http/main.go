package main

import (
	"log"
	"net/http"

	"github.com/joho/godotenv"
	"github.com/kwamekyeimonies/Go-Http/controller"
	"github.com/kwamekyeimonies/Go-Http/database"
)

func main() {

	if err := godotenv.Load(); err !=nil{
		log.Fatal(err)
	}

	db := database.Database_Connection()
	defer db.Close()

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodPost {
			switch r.URL.Path {
			case "/createadmin":
				controller.CreateUser(w,r)
			case "/createproduct":
				controller.CreateProduct(w,r)
			default:
				http.Error(w, "Not Found", http.StatusNotFound)
			}
		} else {
			http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		}
	})
	err := http.ListenAndServe(":4190", nil)
	if err != nil {
		log.Fatal("Listen and Server: ", err)
	}
}
