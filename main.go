package main

import (
	"net/http"

	_ "github.com/go-sql-driver/mysql"

	"github.com/Jhens2020/proyect2020/database"
	"github.com/Jhens2020/proyect2020/persona"
	"github.com/go-chi/chi"
)

func main() {
	db := database.InitDB()

	defer db.Close()
	var personaRepository = persona.NewRepository(db)
	var personaServicio = persona.NewService(personaRepository)

	r := chi.NewRouter()

	//r.Use(helper.GetCors().Handler)*/

	r.Mount("/persona", persona.MakeHTTPSHandler(personaServicio))

	//PORT := os.Getenv("PORT")
	//if PORT == "" {
	//	PORT = "8080"
	//}

	http.ListenAndServe(":3000", r)

}
