package persona

import "database/sql"

/*Repository permite interactuar con la BD*/
type Repository interface {
	InsertPersona(persona *addPersonaRequest) (int, error)
	ObtenerPersonaPorID(param *getPersonaByIDRequest) (*Persona, error)
}

type repository struct {
	db *sql.DB
}

/*NewRepository permite crear el repositoriio*/
func NewRepository(dataBaseConnection *sql.DB) Repository {
	return &repository{
		db: dataBaseConnection,
	}
}
func (repositorysegundo *repository) InsertPersona(persona *addPersonaRequest) (int, error) {
	const queryStr = `insert into PERSONA(NOMBRE,APELLIDO_P,APELLIDO_M,GENERO,DNI,FECHA_NACIMIENTO) values(?,?,?,?,?,?)`
	result, err := repositorysegundo.db.Exec(queryStr, persona.Nombre, persona.ApellidoPat, persona.ApellidoMat, persona.Genero, persona.Dni, persona.FechaNacimiento)
	idpersona, err := result.LastInsertId()
	return int(idpersona), err
}

func (repositorysegundo *repository) ObtenerPersonaPorID(param *getPersonaByIDRequest) (*Persona, error) {
	const queryStr = `SELECT PERSONA_ID, NOMBRE, APELLIDO_P, APELLIDO_M, GENERO, DNI, FECHA_NACIMIENTO FROM PERSONA where PERSONA_ID= ? and ESTADO=1`
	result := repositorysegundo.db.QueryRow(queryStr, param.ID)

	persona := &Persona{}
	err := result.Scan(&persona.ID, &persona.Nombre, &persona.ApellidoPat, &persona.ApellidoMat, &persona.Genero, &persona.DNI, &persona.FechaNacimiento)
	return persona, err
}
