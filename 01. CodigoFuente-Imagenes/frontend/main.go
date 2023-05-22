package main

import (
	"encoding/json"
	"html/template"
	"log"
	"net/http"
	"os"
	"math/rand"
	"time"
)

type Persona struct {
	FotoURL     string `json:"foto_url"`
	Nombre      string `json:"nombre"`
	Cargo       string `json:"cargo"`
	Correo      string `json:"correo"`
	Descripcion string `json:"descripcion"`
}

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		backendURL := os.Getenv("BACKEND_URL")
		resp, err := http.Get(backendURL)
		if err != nil {
			log.Println(err)
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}
		defer resp.Body.Close()

		var personas []Persona
		decoder := json.NewDecoder(resp.Body)
		err = decoder.Decode(&personas)
		if err != nil {
			log.Println(err)
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}

		// Crear un generador de números aleatorios y usarlo para desordenar el slice de personas
		rand.Seed(time.Now().UnixNano())
		rand.Shuffle(len(personas), func(i, j int) { personas[i], personas[j] = personas[j], personas[i] })

		tmpl, err := template.ParseFiles("template.html")
		if err != nil {
			log.Println(err)
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}

		err = tmpl.Execute(w, personas)
		if err != nil {
			log.Println(err)
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}
	})

	log.Println("Listening on :8080...")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal(err)
	}
}
