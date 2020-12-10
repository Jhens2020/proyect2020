package persona

import (
	"context"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/go-chi/chi"
	kithttp "github.com/go-kit/kit/transport/http"
)

/*MakeHTTPSHandler hacer llamados de los metdos*/
func MakeHTTPSHandler(service Service) http.Handler {
	r := chi.NewRouter()

	//Agregar a una persona
	addPersonHandler := kithttp.NewServer(
		makeAddPersonEndpoint(service),
		addPersonRequestDecoder,
		kithttp.EncodeJSONResponse,
	)
	r.Method(http.MethodPost, "/", addPersonHandler)

	//Buscar a una persona por ID
	getPersonaByIDHandler := kithttp.NewServer(
		makeGetPersonaByIDEndPoint(service),
		getPersonaByIDRequestDecoder,
		kithttp.EncodeJSONResponse,
	)
	r.Method(http.MethodGet, "/id/{id}", getPersonaByIDHandler)

	//Buscar a una persona por DNI
	getPersonaByDNIHandler := kithttp.NewServer(
		makeGetPersonaByDNIEndPoint(service),
		getPersonaByDNIRequestDecoder,
		kithttp.EncodeJSONResponse,
	)
	r.Method(http.MethodGet, "/dni/{dni}", getPersonaByDNIHandler)

	return r
}

func addPersonRequestDecoder(context context.Context, r *http.Request) (interface{}, error) {
	request := addPersonaRequest{}
	err := json.NewDecoder(r.Body).Decode(&request)

	return request, err
}

func getPersonaByIDRequestDecoder(context context.Context, r *http.Request) (interface{}, error) {
	personaID, err := strconv.Atoi(chi.URLParam(r, "id"))
	request := getPersonaByIDRequest{
		ID: personaID,
	}
	return request, err
}

func getPersonaByDNIRequestDecoder(context context.Context, r *http.Request) (interface{}, error) {
	personaDNI := (chi.URLParam(r, "dni"))
	request := getPersonaByDNIRequest{
		DNI: personaDNI,
	}
	return request, nil
}
