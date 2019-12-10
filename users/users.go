package users

import (
	"encoding/json"
	"fmt"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	"github.com/solrac97gr/api-gp/database"
)

type User struct {
	Id    int    `json:"id"`
	Email string `json:"email"`
	Name  string `json:"name"`
	Date  string `json:"date"`
}

var err error

func GetUsers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var users []User
	result, err := database.DBCon.Query("SELECT id,email,name,date from users")
	if err != nil {
		panic(err.Error())
	}
	defer result.Close()
	for result.Next() {

		var user User
		err := result.Scan(&user.Id, &user.Email, &user.Name, &user.Date)
		if err != nil {
			panic(err.Error())
		}
		users = append(users, user)
	}
	json.NewEncoder(w).Encode(users)
}
func CreateUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	params := mux.Vars((r))
	_, err := database.DBCon.Query("INSERT INTO users(name,email) VALUES(?,?)", params["name"], params["email"])
	if err != nil {
		panic(err.Error())
	}
	fmt.Fprintf(w, "usuario creado")
}
func GetUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	result, err := database.DBCon.Query("SELECT * FROM users WHERE id = ?", params["id"])
	if err != nil {
		panic(err.Error())
	}
	defer result.Close()
	var user User
	for result.Next() {
		err := result.Scan(&user.Id, &user.Email, &user.Name, &user.Date)
		if err != nil {
			panic(err.Error())
		}
	}
	json.NewEncoder(w).Encode(user)
}
