package persona

import "strconv"

/*Service interface para los sercicios*/
type Service interface {
	InsertPersona(persona *addPersonaRequest) (interface{}, error)
	ObtenerPersonaPorID(param *getPersonaByIDRequest) (*Persona, error)
	ObtenerPersonaPorDNI(param *getPersonaByDNIRequest) (*Persona, error)
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

func (service *service) ObtenerPersonaPorID(param *getPersonaByIDRequest) (*Persona, error) {
	result, err := service.repo.ObtenerPersonaPorID(param)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (service *service) ObtenerPersonaPorDNI(param *getPersonaByDNIRequest) (*Persona, error) {
	result, err := service.repo.ObtenerPersonaPorDNI(param)
	if err != nil {
		return nil, err
	}
	return result, nil
}
