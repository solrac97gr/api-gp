package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
        "flag"
	"net/http"
        "os"
"path/filepath"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	"github.com/jung-kurt/gofpdf"
	"github.com/rs/cors"
)

type User struct {
	Id    int    `json:"id"`
	Email string `json:"email"`
	Name  string `json:"name"`
	Date  string `json:"date"`
}

type Buyer struct {
	id                 string
	name               string
	edad               string
	educacion          string
	redes              string
	industria          string
	n_empleados        string
	canal_comunicacion string
	responsabilidades  string
	superior           string
	aprende_en         string
	herramientas       string
	metrica            string
	objetivos          string
	dificultades       string
}

const (
	STATIC_DIR = "/static/"
)

var db *sql.DB
var err error

func main() {
	db, err = sql.Open("mysql", "apigo:-solrac97G@tcp(167.99.149.193)/api_go")
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

	router := mux.NewRouter()

	router.HandleFunc("/", apiDoc).Methods("GET")
	router.HandleFunc("/users", getUsers).Methods("GET")
	router.HandleFunc("/create/{name}/{email}", createUser).Methods("GET")
	router.HandleFunc("/user/{id}", getUser).Methods("GET")
	router.HandleFunc("/buyer/{user_id}/{name}/{edad}/{education}/{redes}/{industria}/{n_empleados}/{canal_comunicacion}/{responsabilidades}/{superior}/{aprende_en}/{herramientas}/{metrica}/{objetivos}/{dificultades}", buyerPdf).Methods("GET")
	
       var dir string

	flag.StringVar(&dir, "dir", ".", "the directory to serve files from. Defaults to the current dir")
	router.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir(filepath.Dir(os.Args[0])+"/static/"))))

	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"http://localhost:8500","https://carlosgrowth.com","http://localhost:3000"},
		AllowCredentials: false,
	})

	handler := c.Handler(router)
	log.Fatal(http.ListenAndServe(":8500", handler))
}
func getUsers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var users []User
	result, err := db.Query("SELECT id,email,name,date from users")

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
func createUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	params := mux.Vars((r))
	_, err := db.Query("INSERT INTO users(email,name) VALUES(?,?)", params["name"], params["email"])
	if err != nil {
		panic(err.Error())
	}
	fmt.Fprintf(w, "usuario creado")
}
func getUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	result, err := db.Query("SELECT * FROM users WHERE id = ?", params["id"])
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

func buyerPdf(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	var buyer Buyer
	params := mux.Vars(r)

	buyer.id = params["user_id"]
	buyer.name = params["name"]
	buyer.edad = params["edad"]
	buyer.educacion = params["education"]
	buyer.redes = params["redes"]
	buyer.industria = params["industria"]
	buyer.n_empleados = params["n_empleados"]
	buyer.canal_comunicacion = params["canal_comunicacion"]
	buyer.responsabilidades = params["responsabilidades"]
	buyer.superior = params["superior"]
	buyer.aprende_en = params["aprende_en"]
	buyer.herramientas = params["herramientas"]
	buyer.metrica = params["metrica"]
	buyer.objetivos = params["objetivos"]
	buyer.dificultades = params["dificultades"]

	pdf := gofpdf.New("P", "mm", "A4", "")
	pdf.AddPage()
	pdf.ImageOptions(
		filepath.Dir(os.Args[0])+"/static/avatar.png",
		170, 1,
		20, 20,
		false,
		gofpdf.ImageOptions{ImageType: "PNG", ReadDpi: true},
		0,
		"",
	)
	pdf.ImageOptions(
		filepath.Dir(os.Args[0])+"/static/logo.png",
		10, 10,
		40, 10,
		true,
		gofpdf.ImageOptions{ImageType: "PNG", ReadDpi: true},
		0,
		"",
	)
	pdf.SetFont("Arial", "B", 21)
	pdf.SetTextColor(130, 29, 36)
	pdf.CellFormat(190, 7, "", "0", 1, "CM", false, 0, "")
	// CellFormat(width, height, text, border, position after, align, fill, link, linkStr)
	pdf.CellFormat(190, 7, "Buyer Persona", "0", 2, "CM", false, 0, "")
	pdf.SetTextColor(0, 0, 0)

	pdf.SetFont("Arial", "B", 14)
	pdf.SetTextColor(130, 29, 36)
	pdf.MultiCell(190, 7, "Que es:", "0", "L", false)
	pdf.SetFont("Arial", "", 12)
	pdf.SetTextColor(0, 0, 0)

	pdf.MultiCell(190, 7, "Una buyer persona es una representacion semi-ficticia de nuestro consumidor final (o potencial) construida a partir su informacion demografica, comportamiento, necesidades y motivaciones. Al final, se trata de ponernos aun mas en los zapatos de nuestro publico objetivo para entender que necesitan de nosotros.", "0", "L", false)
	pdf.SetFont("Arial", "B", 14)
	pdf.SetTextColor(130, 29, 36)
	pdf.MultiCell(190, 7, "Para que sirve:", "0", "L", false)
	pdf.SetFont("Arial", "", 12)
	pdf.SetTextColor(0, 0, 0)

	pdf.MultiCell(190, 7, "Los buyer personas son tan importantes hoy en dia porque ayudan a entender mejor a los clientes actuales y potenciales. Ademas, es importante tener en cuenta que facilitan la creacian y planificacion de contenido relevante y permiten saber como hay que desarrollar los productos y que tipo de servicios ofrecer dependiendo de sus comportamientos, necesidades y preocupaciones. En definitiva, el buyer persona nos permite disenar acciones de marketing mas efectivas.", "0", "L", false)
	pdf.CellFormat(150, 7, "", "0", 3, "L", false, 0, "")
	pdf.SetFont("Arial", "B", 12)
	pdf.SetTextColor(255, 255, 255)
	pdf.SetFillColor(130, 29, 36)
	pdf.MultiCell(150, 7, "Esta es tu persona", "1", "L", true)
	pdf.SetFont("Arial", "", 12)
	pdf.SetTextColor(0, 0, 0)
	pdf.CellFormat(150, 7, "Se llama: "+buyer.name, "1", 3, "L", false, 0, "")
	pdf.CellFormat(150, 7, "Tiene: "+buyer.edad, "1", 4, "L", false, 0, "")
	pdf.CellFormat(150, 7, "Redes que usa: "+buyer.redes, "1", 5, "L", false, 0, "")
	pdf.CellFormat(150, 7, "Industria en la que esta: "+buyer.industria, "1", 6, "L", false, 0, "")
	pdf.CellFormat(150, 7, "Numero de empleados: "+buyer.n_empleados, "1", 7, "L", false, 0, "")
	pdf.CellFormat(150, 7, "Se comunica usando: "+buyer.canal_comunicacion, "1", 8, "L", false, 0, "")
	pdf.CellFormat(150, 7, "Sus responsabilidades son: "+buyer.responsabilidades, "1", 9, "L", false, 0, "")
	pdf.CellFormat(150, 7, "Su Jefe es : "+buyer.superior, "1", 10, "L", false, 0, "")
	pdf.CellFormat(150, 7, "Aprende usando : "+buyer.aprende_en, "1", 11, "L", false, 0, "")
	pdf.CellFormat(150, 7, "Usa estas herramientas: "+buyer.herramientas, "1", 12, "L", false, 0, "")
	pdf.CellFormat(150, 7, "Su Metrica: "+buyer.metrica, "1", 13, "L", false, 0, "")
	pdf.CellFormat(150, 7, "Sus Objetivos: "+buyer.objetivos, "1", 14, "L", false, 0, "")
	pdf.CellFormat(150, 7, "Sus Dificultades: "+buyer.dificultades, "1", 15, "L", false, 0, "")

	pdf.OutputFileAndClose(filepath.Dir(os.Args[0])+"/static/buyer_" + params["user_id"] + ".pdf")

	if err != nil {
		panic(err)
	}
var ruta = os.Args[0]
	fmt.Fprintf(w, `Done`+ruta)
}
func apiDoc(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, `Bienvenidos al api`)
}

