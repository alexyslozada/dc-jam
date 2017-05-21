package handlers

import (
	"net/http"
	"io/ioutil"
	"fmt"
	"strings"
)

const datosAbiertosURL = "http://datosabiertos.vivelabbogota.com/api/action/datastore_search"

type Message struct {
	Code    int
	Message string
	Error   string
}

func showConsole(r *http.Response) {
	body, _ := ioutil.ReadAll(r.Body)
	fmt.Println("response Body:", string(body))
}

func removeDot(s string) string {
	s = strings.Replace(s, ".", "", -1)
	return s
}

