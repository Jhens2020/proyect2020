package persona

import "strconv"

/*Service interface para los sercicios*/
type Service interface {
	InsertPersona(persona *addPersonaRequest) (interface{}, error)
}

type service struct {
	repo Repository
}

/*NewService permite manipular los metodos del repositorio*/
func NewService(repo Repository) Service {
	return &service{
		repo: repo,
	}
}
func (service *service) InsertPersona(persona *addPersonaRequest) (interface{}, error) {
	result, err := service.repo.InsertPersona(persona)
	if err != nil {
		return nil, err
	}
	mensaje := "Persona registrada con el id " + strconv.Itoa(result)

	return map[string]string{"mensaje": mensaje}, nil
}
