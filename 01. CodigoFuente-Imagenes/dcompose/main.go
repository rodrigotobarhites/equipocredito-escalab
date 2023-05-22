package main

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"
	"os"

	_ "github.com/lib/pq"
)

// Definición de la estructura para almacenar cada fila de la tabla
type ActividadEstadistica struct {
	NombreBD string `json:"nombre_bd"`
	Estado   string `json:"estado"`
}

func main() {
	// Definición del manejador para la ruta '/'
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		// Conexión a la base de datos PostgreSQL
		bdd, err := sql.Open("postgres", "host="+os.Getenv("DB_HOST")+" user="+os.Getenv("DB_USER")+" password="+os.Getenv("DB_PASSWORD")+" dbname=postgres sslmode=disable")
		if err != nil {
			log.Fatal(err)
		}

		// Consulta para seleccionar todas las filas de la tabla 'pg_stat_activity' que no tienen NULL en la columna 'state'
		fila, err := bdd.Query("SELECT datname, state FROM pg_stat_activity WHERE state IS NOT NULL")
		if err != nil {
			log.Fatal(err)
		}

		// Definición del slice para almacenar todas las filas de la tabla
		var actividades []ActividadEstadistica

		// Iteración por cada fila de la tabla
		for fila.Next() {
			var actividad ActividadEstadistica
			if err := fila.Scan(&actividad.NombreBD, &actividad.Estado); err != nil {
				log.Fatal(err)
			}
			actividades = append(actividades, actividad)
		}

		// Conversión de las filas a JSON y escritura en la respuesta HTTP
		json.NewEncoder(w).Encode(actividades)
	})

	// Inicio del servidor HTTP en el puerto 8080
	log.Fatal(http.ListenAndServe(":8080", nil))
}





