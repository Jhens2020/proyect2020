package persona

import (
	"context"

	"github.com/go-kit/kit/endpoint"
)

type addPersonaRequest struct {
	Nombre          string
	ApellidoPat     string
	ApellidoMat     string
	Genero          string
	Dni             string
	FechaNacimiento string
}
type getPersonaByIDRequest struct {
	ID int
}
type getPersonaByDNIRequest struct {
	DNI string
}

/* Permite mostrar json body*/
func makeAddPersonEndpoint(service Service) endpoint.Endpoint {
	addPersonEndpoint := func(ctx context.Context, request interface{}) (interface{}, error) {
		// req := request.(addPersonaRequest)
		// persona_id, err := s.InsertPerson(&req)
		rep := request.(addPersonaRequest)
		result, err := service.InsertPersona(&rep)

		return result, err
	}
	return addPersonEndpoint
}

func makeGetPersonaByIDEndPoint(service Service) endpoint.Endpoint {
	getPersonaByID := func(ctx context.Context, request interface{}) (interface{}, error) {
		// req := request.(addPersonaRequest)
		// persona_id, err := s.InsertPerson(&req)
		rep := request.(getPersonaByIDRequest)
		result, err := service.ObtenerPersonaPorID(&rep)

		return result, err
	}
	return getPersonaByID
}

func makeGetPersonaByDNIEndPoint(service Service) endpoint.Endpoint {
	GetPersonaByDNI := func(ctx context.Context, request interface{}) (interface{}, error) {
		// req := request.(addPersonaRequest)
		// persona_id, err := s.InsertPerson(&req)
		rep := request.(getPersonaByDNIRequest)
		result, err := service.ObtenerPersonaPorDNI(&rep)

		return result, err
	}
	return GetPersonaByDNI
}
