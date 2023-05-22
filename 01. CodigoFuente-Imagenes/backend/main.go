package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

type Persona struct {
	FotoURL     string `json:"foto_url"`
	Nombre      string `json:"nombre"`
	Cargo       string `json:"cargo"`
	Correo      string `json:"correo"`
	Descripcion string `json:"descripcion"`
}

var personas []Persona

func cargarDatos() {
	data, err := ioutil.ReadFile("datos.json")
	if err != nil {
		log.Fatalf("No se pudo leer el archivo: %v", err)
	}

	err = json.Unmarshal(data, &personas)
	if err != nil {
		log.Fatalf("No se pudo deserializar los datos: %v", err)
	}
}

func obtenerIntegrantes(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(personas)
}

func liveness(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}

func main() {
	cargarDatos()

	http.HandleFunc("/obtenerIntegrantes", obtenerIntegrantes)
	http.HandleFunc("/liveness", liveness)

	log.Fatal(http.ListenAndServe(":8080", nil))
}
