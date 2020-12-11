package main

import (
	"net/http"
	"os"

	_ "github.com/go-sql-driver/mysql"

	"github.com/Jhens2020/proyect2020/database"
	"github.com/Jhens2020/proyect2020/handlers"
	"github.com/go-chi/chi"
)

func main() {
	db := database.InitDB()

	defer db.Close()
	r := chi.NewRouter()
	r.Mount("/v1", handlers.MakeHTTPSHandlerV1(db))
	r.Mount("/v2", handlers.MakeHTTPSHandlerV2(db))

	PORT := os.Getenv("PORT")
	if PORT == "" {
		PORT = "8080"
	}

	http.ListenAndServe(":"+PORT, r)

}
