package handlers

import (
	"database/sql"
	"net/http"

	"github.com/Jhens2020/proyect2020/persona"
	"github.com/go-chi/chi"
)

/*MakeHTTPSHandlerV2 sirve para tener el informe de Versiones de la ruta*/
func MakeHTTPSHandlerV2(db *sql.DB) http.Handler {
	r := chi.NewRouter()

	var personaRepository = persona.NewRepository(db)
	var personaServicio = persona.NewService(personaRepository)

	//r.Use(helper.GetCors().Handler)*/

	r.Mount("/persona", persona.MakeHTTPSHandler(personaServicio))

	return r
}
