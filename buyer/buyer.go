package buyer

import (
	"encoding/json"
	"net/http"
	"strconv"
)

var err error

type Buyer struct {
	//PERSONAL
	Id                int     `json:"id"`
	Nombre            string  `json:"nombre"`
	Edad              int     `json:"edad"`
	Estado_civil      string  `json:"estado_civil"`
	Educacion         string  `json:"educacion"`
	Condicion_laboral string  `json:"condicion_laboral"`
	Titulo            string  `json:"titulo"`
	Ingresos_anuales  float64 `json:"ingresos_anuales"`
	//CONDUCTA PERSONAL
	Necesita_ser_feliz     string `json:"necesita_ser_feliz"`
	Que_hace_no_trabajando string `json:"que_hace_no_trabajando"`
	En_que_gasta           string `json:"en_que_gasta"`
	Donde_pasa_mas_tiempo  string `json:"donde_pasa_mas_tiempo"`
	Como_mide_exito        string `json:"como_mide_exito"`
	Personas_importantes   string `json:"personas_importantes"`
	//CONDUCTA ONLINE
	Tiempo_internet              string `json:"tiempo_internet"`
	Dispositivos_usa             string `json:"dispositivos_usa"`
	Red_social_preferida         string `json:"red_social_preferida"`
	Blogs_favoritos              string `json:"blogs_favoritos"`
	Contenido_lee                string `json:"contenido_lee"`
	Mayores_intereses            string `json:"mayores_intereses"`
	Donde_busca_info             string `json:"donde_busca_info"`
	Formato_prefiere_aprender    string `json:"formato_prefiere_aprender"`
	Principal_actividad_internet string `json:"principal_actividad_internet"`
	Principal_informacion_busca  string `json:"principal_informacion_busca"`
	Marcas_sigue                 string `json:"marcas_sigue"`
	Compra_online                string `json:"compra_online"`
	Horario_internet             string `json:"horario_internet"`
	Influenciadores_online       string `json:"influenciadores_online"`
	Lenguaje_usado               string `json:"lenguaje_usado"`

	//CONDUCTA LABORAL
	Problema_laboral_solucionar string `json:"problema_laboral_solucionar"`
	Mayor_responsabilidad       string `json:"mayor_responsabilidad"`
	Problema_laboral_infeliz    string `json:"problema_laboral_infeliz"`
	Aptitudes_buen_trabajo      string `json:"aptitudes_buen_trabajo"`
	Superior                    string `json:"superior"`
	Influencia_laboral          string `json:"influencia_laboral"`
	Aspiracion_laboral          string `json:"aspiracion_laboral"`

	//RELACIÓN CON LA EMPRESA
	Porque_nos_necesita          string `json:"porque_nos_necesita"`
	Como_nos_conocieron          string `json:"como_nos_conocieron"`
	Aspectos_ayudamos            string `json:"aspectos_ayudamos"`
	Aspecto_evalua_antes_comprar string `json:"aspecto_evalua_antes_comprar"`
	Sentimiento_abandonarlo      string `json:"sentimiento_abandonarlo"`
	Impacto_empresa_vida         string `json:"impacto_empresa_vida"`
	Impacto_financiero           string `json:"impacto_financiero"`
	Mayor_objecion               string `json:"mayor_objecion"`
	Tiempo_colaborando           string `json:"tiempo_colaborando"`
}

func BuyerPdf(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	var persona Buyer
	//DATA PERSONAL
	persona.Nombre = r.FormValue("nombre")
	persona.Edad, err = strconv.Atoi(r.FormValue("edad"))
	if err != nil {
		panic(err.Error())
	}
	persona.Estado_civil = r.FormValue("estado_civil")
	persona.Educacion = r.FormValue("educacion")
	persona.Condicion_laboral = r.FormValue("nombre")
	persona.Titulo = r.FormValue("titulo")
	persona.Ingresos_anuales, err = strconv.ParseFloat(r.FormValue("ingresos_anuales"), 32)
	if err != nil {
		panic(err.Error())
	}
	//CONDUCTA PERSONAL
	persona.Necesita_ser_feliz = r.FormValue("necesita_ser_feliz")
	persona.Que_hace_no_trabajando = r.FormValue("que_hace_no_trabajando")
	persona.En_que_gasta = r.FormValue("en_que_gasta")
	persona.Donde_pasa_mas_tiempo = r.FormValue("donde_pasa_mas_tiempo")
	persona.Como_mide_exito = r.FormValue("como_mide_exito")
	persona.Personas_importantes = r.FormValue("personas_importantes")
	//CONDUCTA ONLINE
	persona.Tiempo_internet = r.FormValue("tiempo_internet")
	persona.Dispositivos_usa = r.FormValue("dispositivos_usa")
	persona.Red_social_preferida = r.FormValue("red_social_preferida")
	persona.Blogs_favoritos = r.FormValue("blogs_favoritos")
	persona.Contenido_lee = r.FormValue("contenido_lee")
	persona.Mayores_intereses = r.FormValue("mayores_intereses")
	persona.Donde_busca_info = r.FormValue("donde_busca_info")
	persona.Formato_prefiere_aprender = r.FormValue("formato_prefiere_aprender")
	persona.Principal_actividad_internet = r.FormValue("principal_actividad_internet")
	persona.Principal_informacion_busca = r.FormValue("principal_informacion_busca")
	persona.Marcas_sigue = r.FormValue("marcas_sigue")
	persona.Compra_online = r.FormValue("compra_online")
	persona.Horario_internet = r.FormValue("horario_internet")
	persona.Influenciadores_online = r.FormValue("influenciadores_online")
	persona.Lenguaje_usado = r.FormValue("lenguaje_usado")
	//CONDUCTA LABORAL
	persona.Problema_laboral_solucionar = r.FormValue("problema_laboral_solucionar")
	persona.Mayor_responsabilidad = r.FormValue("mayor_responsabilidad")
	persona.Problema_laboral_infeliz = r.FormValue("problema_laboral_infeliz")
	persona.Aptitudes_buen_trabajo = r.FormValue("aptitudes_buen_trabajo")
	persona.Superior = r.FormValue("superior")
	persona.Influencia_laboral = r.FormValue("influencia_laboral")
	persona.Aspiracion_laboral = r.FormValue("aspiracion_laboral")

	//RELACIÓN CON LA EMPRESA
	persona.Porque_nos_necesita = r.FormValue("porque_nos_necesita")
	persona.Como_nos_conocieron = r.FormValue("como_nos_conocieron")
	persona.Aspectos_ayudamos = r.FormValue("aspectos_ayudamos")
	persona.Aspecto_evalua_antes_comprar = r.FormValue("aspecto_evalua_antes_comprar")
	persona.Sentimiento_abandonarlo = r.FormValue("sentimiento_abandonarlo")
	persona.Impacto_empresa_vida = r.FormValue("impacto_empresa_vida")
	persona.Impacto_financiero = r.FormValue("impacto_financiero")
	persona.Mayor_objecion = r.FormValue("mayor_objecion")
	persona.Tiempo_colaborando = r.FormValue("tiempo_colaborando")
	//RESPONDE CON LO QUE LE MANDAN
	json.NewEncoder(w).Encode(&persona)
}

/*
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

} */
