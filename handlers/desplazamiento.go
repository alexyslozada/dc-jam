package handlers

import (
	"fmt"
	"github.com/labstack/echo"
	"github.com/labstack/gommon/log"
	"net/http"
	"encoding/json"
	"strconv"
)

type DesplazamientoStruct struct {
	ID                 int    `json:"_id"`
	PersonasExpulsadas string `json:"PERSONAS EXPULSADAS"`
	PersonasDeclaradas string `json:"PERSONAS DECLARADAS"`
	Vigencia           string `json:"VIGENCIA"`
	PersonasRecibidas  string `json:"PERSONAS RECIBIDAS"`
	Expulsadas         int    `json:"expulsadas"`
	Declaradas         int    `json:"declaradas"`
	Recibidas          int    `json:"recibidas"`
}

type ResponseDesplazamiento struct {
	Help string `json:"help"`
	Success bool `json:"success"`
	Result struct {
		ResourceID string `json:"resource_id"`
		Fields []struct {
			Type string `json:"type"`
			ID string `json:"id"`
		}
		Records []*DesplazamientoStruct `json:"records"`
		Links struct {
			Start string `json:"start"`
			Next string `json:"next"`
		} `json:"_links"`
		Total int `json:"total"`
	} `json:"result"`
}

func Desplazamiento(c echo.Context) error {
	m := Message{}
	rd := ResponseDesplazamiento{}
	resourceID := "9d255c81-07f9-4d3d-93ed-5c02ad7ad146"
	strURL := fmt.Sprintf("%s?resource_id=%s", datosAbiertosURL, resourceID)

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
	err = json.NewDecoder(resp.Body).Decode(&rd)
	if err != nil {
		m.Code = http.StatusInternalServerError
		m.Message = "No fue posible convertir el objeto recibido"
		m.Error = err.Error()
		log.Print(err)
		return c.JSON(m.Code, m)
	}
	err = setValuesDesplazamiento(rd.Result.Records)
	if err != nil {
		m.Code = http.StatusInternalServerError
		m.Message = "No se pudo convertir los valores enteros"
		m.Error = err.Error()
		return c.JSON(m.Code, m)
	}

	return c.JSON(http.StatusOK, rd)
}

func setValuesDesplazamiento(rd []*DesplazamientoStruct) error {
	for _, r := range rd {
		vd, err := strconv.Atoi(removeDot(r.PersonasDeclaradas))
		if err != nil {
			return err
		}
		ve, err := strconv.Atoi(removeDot(r.PersonasExpulsadas))
		if err != nil {
			return err
		}
		vr, err := strconv.Atoi(removeDot(r.PersonasRecibidas))
		if err != nil {
			return err
		}

		r.Declaradas = vd
		r.Expulsadas = ve
		r.Recibidas = vr
	}
	return nil
}
