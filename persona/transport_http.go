package persona

import (
	"context"
	"encoding/json"
	"net/http"

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

	return r
}

func addPersonRequestDecoder(context context.Context, r *http.Request) (interface{}, error) {
	request := addPersonaRequest{}
	err := json.NewDecoder(r.Body).Decode(&request)

	return request, err
}
