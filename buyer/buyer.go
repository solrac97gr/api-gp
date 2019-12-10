package buyer

import (
	"net/http"
	"os"
	"path/filepath"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/jung-kurt/gofpdf"
)

type Buyer struct {
	//PERSONAL
	id                int
	nombre            string
	edad              int
	estado_civil      string
	educacion         string
	condicion_laboral string
	titulo            string
	ingresos_anuales  float32

	necesita_ser_feliz     string
	que_hace_no_trabajando string
	en_que_gasta           string
	donde_pasa_mas_tiempo  string
	como_mide_exito        string
	personas_importantes   string
	//CONDUCTA ONLINE
	tiempo_internet              string
	dispositivos_usa             string
	red_social_preferida         string
	blogs_favoritos              string
	contenido_lee                string
	mayores_intereses            string
	donde_busca_info             string
	formato_prefiere_aprender    string
	principal_actividad_internet string
	principal_informacion_busca  string
	marcas_sigue                 string
	compra_online                string
}

func BuyerPdf(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	var buyer Buyer
	params := mux.Vars(r)
	convert_id, err := strconv.Atoi(params["id"])
	if err != nil {
		panic(err.Error())
	}

	conver_edad, err := strconv.Atoi(params["edad"])
	if err != nil {
		panic(err.Error())
	}
	conver_empleados, err := strconv.Atoi(params["n_empleados"])
	if err != nil {
		panic(err.Error())
	}
	conver_typeid, err := strconv.Atoi(params["type_id"])

	//ints converted
	buyer.id = convert_id
	buyer.edad = conver_edad
	buyer.n_empleados = conver_empleados
	buyer.type_id = conver_typeid

	//string data

	buyer.name = params["name"]
	buyer.educacion = params["education"]
	buyer.redes = params["redes"]
	buyer.industria = params["industria"]
	buyer.canal_comunicacion = params["canal_comunicacion"]
	buyer.responsabilidades = params["responsabilidades"]
	buyer.superior = params["superior"]
	buyer.aprende_en = params["aprende_en"]
	buyer.herramientas = params["herramientas"]
	buyer.metrica = params["metrica"]
	buyer.objetivos = params["objetivos"]
	buyer.dificultades = params["dificultades"]

	if conver_typeid == 1 {
		generatePDFB2B(buyer)
	} else {
		generatePDFB2C(buyer)
	}

}

func generatePDFB2B(buyer Buyer) {
	//convert INT to string
	convert_id := strconv.Itoa(buyer.id)
	convert_edad := strconv.Itoa(buyer.edad)
	convert_nempleados := strconv.Itoa(buyer.n_empleados)

	pdf := gofpdf.New("P", "mm", "A4", "")
	tr := pdf.UnicodeTranslatorFromDescriptor("")
	pdf.AddPage()
	pdf.ImageOptions(
		filepath.Dir(os.Args[0])+"/static/brand/avatar.png",
		170, 1,
		20, 20,
		false,
		gofpdf.ImageOptions{ImageType: "PNG", ReadDpi: true},
		0,
		"",
	)
	pdf.ImageOptions(
		filepath.Dir(os.Args[0])+"/static/brand/logo.png",
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
	pdf.MultiCell(190, 7, tr("¿Qué es?:"), "0", "L", false)
	pdf.SetFont("Arial", "", 12)
	pdf.SetTextColor(0, 0, 0)

	pdf.MultiCell(190, 7, tr("Una buyer persona es una representación semi-ficticia de nuestro consumidor final (o potencial) construida a partir su información demográfica, comportamiento, necesidades y motivaciones. Al final, se trata de ponernos aún más en los zapatos de nuestro público objetivo para entender qué necesitan de nosotros."), "0", "L", false)
	pdf.SetFont("Arial", "B", 14)
	pdf.SetTextColor(130, 29, 36)
	pdf.MultiCell(190, 7, tr("¿Para qué sirve?:"), "0", "L", false)
	pdf.SetFont("Arial", "", 12)
	pdf.SetTextColor(0, 0, 0)

	pdf.MultiCell(190, 7, tr("Los buyer personas son tan importantes hoy en día porque ayudan a entender mejor a los clientes actuales y potenciales. Además, es importante tener en cuenta que facilitan la creación y planificación de contenido relevante y permiten saber cómo hay que desarrollar los productos y qué tipo de servicios ofrecer dependiendo de sus comportamientos, necesidades y preocupaciones. En definitiva, el buyer persona nos permite diseñar acciones de marketing más efectivas."), "0", "L", false)
	pdf.CellFormat(150, 7, "", "0", 3, "L", false, 0, "")
	pdf.SetFont("Arial", "B", 12)
	pdf.SetTextColor(255, 255, 255)
	pdf.SetFillColor(130, 29, 36)
	pdf.MultiCell(150, 7, tr("Esta es tu persona"), "1", "L", true)
	pdf.SetFont("Arial", "", 12)
	pdf.SetTextColor(0, 0, 0)
	pdf.CellFormat(150, 7, tr("Se llama: "+buyer.name), "1", 3, "L", false, 0, "")
	pdf.CellFormat(150, 7, tr("Tiene: "+convert_edad), "1", 4, "L", false, 0, "")
	pdf.CellFormat(150, 7, tr("Redes que usa: "+buyer.redes), "1", 5, "L", false, 0, "")
	pdf.CellFormat(150, 7, tr("Industria en la que está: "+buyer.industria), "1", 6, "L", false, 0, "")
	pdf.CellFormat(150, 7, tr("Número de empleados: "+convert_nempleados), "1", 7, "L", false, 0, "")
	pdf.CellFormat(150, 7, tr("Se comunica usando: "+buyer.canal_comunicacion), "1", 8, "L", false, 0, "")
	pdf.CellFormat(150, 7, tr("Sus responsabilidades son: "+buyer.responsabilidades), "1", 9, "L", false, 0, "")
	pdf.CellFormat(150, 7, tr("Su Jefe es : "+buyer.superior), "1", 10, "L", false, 0, "")
	pdf.CellFormat(150, 7, tr("Aprende usando : "+buyer.aprende_en), "1", 11, "L", false, 0, "")
	pdf.CellFormat(150, 7, tr("Usa estas herramientas: "+buyer.herramientas), "1", 12, "L", false, 0, "")
	pdf.CellFormat(150, 7, tr("Su Métrica: "+buyer.metrica), "1", 13, "L", false, 0, "")
	pdf.CellFormat(150, 7, tr("Sus Objetivos: "+buyer.objetivos), "1", 14, "L", false, 0, "")
	pdf.CellFormat(150, 7, tr("Sus Dificultades: "+buyer.dificultades), "1", 15, "L", false, 0, "")

	pdf.OutputFileAndClose(filepath.Dir(os.Args[0]) + "/static/buyer_" + convert_id + ".pdf")
}

func generatePDFB2C(buyer Buyer) {

}
