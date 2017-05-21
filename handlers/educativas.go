package handlers

import (
	"github.com/labstack/echo"
	"fmt"
	"net/http"
	"github.com/labstack/gommon/log"
	"encoding/json"
)

func init()  {
	loadLocalidades()
}

type EducativasStruct struct {
	ID int `json:"_id"`
	Sector string `json:"sector"`
	NombreInstitucion string `json:"nombreinstitucion"`
	Localidad string `json:"localidad"`
	Formal string `json:"formal"`
}

type LocalidadStruct struct {
	ID string `json:"id"`
	Nombre string `json:"nombre"`
	Superficie float32 `json:"superficie"`
	Poblacion int `json:"poblacion"`
	Densidad float32 `json:"densidad"`
	Latitud float32 `json:"latitud"`
	Longitud float32 `json:"longitud"`
}

type ResponseEducativas struct {
	Help string `json:"help"`
	Success bool `json:"success"`
	Result struct {
		ResourceID string `json:"resource_id"`
		Fields []struct {
			Type string `json:"type"`
			ID string `json:"id"`
		}
		Records []*EducativasStruct `json:"records"`
		Links struct {
			Start string `json:"start"`
			Next string `json:"next"`
		} `json:"_links"`
		Total int `json:"total"`
	} `json:"result"`
	Aggregate struct {
		Formales map[string]int `json:"formales"`
		NoFormales map[string]int `json:"no_formales"`
	} `json:"aggregate"`
	Localidades []LocalidadStruct `json:"localidades"`
}

var localidades = make([]LocalidadStruct, 0)

func Educativas(c echo.Context) error {
	m := Message{}
	re := ResponseEducativas{}

	resourceID := "92d26b06-96ec-4ac5-9111-be81fa0112b2"
	limit := "10000"
	strURL := fmt.Sprintf("%s?resource_id=%s&limit=%s", datosAbiertosURL, resourceID, limit)

	req, err := http.NewRequest("GET", strURL, nil)
	if err != nil {
		m.Code = http.StatusInternalServerError
		m.Message = "Error al consultar la API"
		m.Error = err.Error()
		log.Print(err)
		return c.JSON(m.Code, m)
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		m.Code = http.StatusInternalServerError
		m.Message = "Error al ejecutar la petición"
		m.Error = err.Error()
		log.Print(err)
		return c.JSON(m.Code, m)
	}
	defer resp.Body.Close()

	// showConsole(resp)
	err = json.NewDecoder(resp.Body).Decode(&re)
	if err != nil {
		m.Code = http.StatusInternalServerError
		m.Message = "No fue posible convertir el objeto recibido"
		m.Error = err.Error()
		log.Print(err)
		return c.JSON(m.Code, m)
	}

	re.Aggregate.Formales = make(map[string]int)
	re.Aggregate.NoFormales = make(map[string]int)

	for _, v := range re.Result.Records {
		if v.Formal == "1" {
			re.Aggregate.Formales[v.Localidad]++
		} else {
			re.Aggregate.NoFormales[v.Localidad]++
		}
	}

	re.Localidades = localidades

	return c.JSON(http.StatusOK, re)
}

func loadLocalidades() {
	localidades = append(localidades, LocalidadStruct{ID: "1", Nombre: "Usaquén", Superficie: 65.31, Poblacion: 501999, Densidad: 7686.4, Latitud: 4.744984, Longitud: -74.0988376})
	localidades = append(localidades, LocalidadStruct{ID: "2", Nombre: "Chapinero", Superficie: 38.15, Poblacion: 139701, Densidad: 3661.88, Latitud: 4.6486972, Longitud: -74.0652175})
	localidades = append(localidades, LocalidadStruct{ID: "3", Nombre: "Santa Fe", Superficie: 45.17, Poblacion: 110048, Densidad: 2436.3, Latitud: 4.5929783, Longitud: -74.0581677})
	localidades = append(localidades, LocalidadStruct{ID: "4", Nombre: "San Cristóbal", Superficie: 49.09, Poblacion: 404697, Densidad: 8243.98, Latitud: 4.5489797, Longitud: -74.1015474})
	localidades = append(localidades, LocalidadStruct{ID: "5", Nombre: "Usme", Superficie: 215.06, Poblacion: 457302, Densidad: 2126.39, Latitud: 4.5029527, Longitud: -74.1302869})
	localidades = append(localidades, LocalidadStruct{ID: "6", Nombre: "Tunjuelito", Superficie: 9.91, Poblacion: 199430, Densidad: 20124.11, Latitud: 4.5637033, Longitud: -74.1558746})
	localidades = append(localidades, LocalidadStruct{ID: "7", Nombre: "Bosa", Superficie: 23.93, Poblacion: 673077, Densidad: 28126.91, Latitud: 4.6256413, Longitud: -74.2051251})
	localidades = append(localidades, LocalidadStruct{ID: "8", Nombre: "Kennedy", Superficie: 38.59, Poblacion: 1088443, Densidad: 28205.31, Latitud: 4.6299452, Longitud: -74.1693876})
	localidades = append(localidades, LocalidadStruct{ID: "9", Nombre: "Fontibón", Superficie: 33.28, Poblacion: 394648, Densidad: 11858.41, Latitud: 4.6775246, Longitud: -74.1747025})
	localidades = append(localidades, LocalidadStruct{ID: "10", Nombre: "Engativá", Superficie: 35.88, Poblacion: 887080, Densidad: 24723.52, Latitud: 4.6971531, Longitud: -74.1536804})
	localidades = append(localidades, LocalidadStruct{ID: "11", Nombre: "Suba", Superficie: 100.56, Poblacion: 1218513, Densidad: 12117.27, Latitud: 4.761433, Longitud: -74.1530326})
	localidades = append(localidades, LocalidadStruct{ID: "12", Nombre: "Barrios Unidos", Superficie: 11.9, Poblacion: 243465, Densidad: 20459.24, Latitud: 4.6697582, Longitud: -74.0928956})
	localidades = append(localidades, LocalidadStruct{ID: "13", Nombre: "Teusaquillo", Superficie: 14.19, Poblacion: 153025, Densidad: 10784, Latitud: 4.6408003, Longitud: -74.1050016})
	localidades = append(localidades, LocalidadStruct{ID: "14", Nombre: "Los Mártires", Superficie: 6.51, Poblacion: 99119, Densidad: 15225.65, Latitud: 4.6084199, Longitud: -74.0987183})
	localidades = append(localidades, LocalidadStruct{ID: "15", Nombre: "Antonio Nariño", Superficie: 4.88, Poblacion: 109176, Densidad: 22372.12, Latitud: 4.5872224, Longitud: -74.1169983})
	localidades = append(localidades, LocalidadStruct{ID: "16", Nombre: "Puente Aranda", Superficie: 17.31, Poblacion: 258287, Densidad: 14921.25, Latitud: 4.6194713, Longitud: -74.1279636})
	localidades = append(localidades, LocalidadStruct{ID: "17", Nombre: "La Candelaria", Superficie: 2.06, Poblacion: 24088, Densidad: 11693.2, Latitud: 4.596399, Longitud: -74.0754114})
	localidades = append(localidades, LocalidadStruct{ID: "18", Nombre: "Rafael Uribe Uribe", Superficie: 13.83, Poblacion: 374246, Densidad: 27060.44, Latitud: 4.5783005, Longitud: -74.1342029})
	localidades = append(localidades, LocalidadStruct{ID: "19", Nombre: "Ciudad Bolívar", Superficie: 130, Poblacion: 707569, Densidad: 5442.83, Latitud: 4.5344827, Longitud: -74.188965})
	localidades = append(localidades, LocalidadStruct{ID: "20", Nombre: "Sumapaz", Superficie: 780.96, Poblacion: 6531, Densidad: 8.36, Latitud: 4.6089062, Longitud: -74.1028789})
}
