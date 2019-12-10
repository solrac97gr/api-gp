package main

import (
	"database/sql"
	"flag"
	"log"
	"net/http"
	"os"
	"path/filepath"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	"github.com/rs/cors"
	"github.com/solrac97gr/api-gp/buyer"
	"github.com/solrac97gr/api-gp/database"
	"github.com/solrac97gr/api-gp/users"
)

var err error

func main() {
	database.DBCon, err = sql.Open("mysql", "apigo:-solrac97G@tcp(167.99.149.193)/api_go")
	if err != nil {
		panic(err.Error())
	}
	defer database.DBCon.Close()
	router := mux.NewRouter()
	router.HandleFunc("/", apiDoc).Methods("GET")
	router.HandleFunc("/users", users.GetUsers).Methods("GET")
	router.HandleFunc("/create/{name}/{email}", users.CreateUser).Methods("GET")
	router.HandleFunc("/user/{id}", users.GetUser).Methods("GET")
	router.HandleFunc("/buyer/{id}/{name}/{edad}/{education}/{redes}/{industria}/{n_empleados}/{canal_comunicacion}/{responsabilidades}/{superior}/{aprende_en}/{herramientas}/{metrica}/{objetivos}/{dificultades}/{type_id}", buyer.BuyerPdf).Methods("GET")

	var dir string

	flag.StringVar(&dir, "dir", ".", "the directory to serve files from. Defaults to the current dir")
	router.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir(filepath.Dir(os.Args[0])+"/static/"))))

	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"http://localhost:8500", "https://carlosgrowth.com", "http://localhost:3000"},
		AllowCredentials: false,
	})

	handler := c.Handler(router)
	log.Fatal(http.ListenAndServe(":8500", handler))
}

func apiDoc(w http.ResponseWriter, r *http.Request) {
	url := filepath.Dir(os.Args[0]) + "/index.html"
	w.Header().Set("Content-type", "text/html")
	http.ServeFile(w, r, url)
}
