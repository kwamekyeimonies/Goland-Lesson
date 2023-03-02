package models

type User struct{
	ID string `json:"id"`
	Email string `json:"email"`
	Password string `jsons:"password"`
}


type Product struct{
	ID string `json:"id"`
	Name string `json:"name"`
}