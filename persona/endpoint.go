package persona

import (
	"context"
	"fmt"

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

/* Permite mostrar json body*/
func makeAddPersonEndpoint(service Service) endpoint.Endpoint {
	addPersonEndpoint := func(ctx context.Context, request interface{}) (interface{}, error) {
		// req := request.(addPersonaRequest)
		// persona_id, err := s.InsertPerson(&req)
		rep := request.(addPersonaRequest)
		result, err := service.InsertPersona(&rep)

		fmt.Println(rep.Nombre)
		fmt.Println(rep.ApellidoPat)
		fmt.Println(rep.ApellidoMat)
		fmt.Println(rep.Genero)
		fmt.Println(rep.Dni)
		fmt.Println(rep.FechaNacimiento)

		return result, err
	}
	return addPersonEndpoint
}
