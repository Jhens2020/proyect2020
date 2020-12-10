package persona

/*Persona esta permite mostra la data de la persona en formato JSONS*/
type Persona struct {
	ID              int    `json:"id"`
	Nombre          string `json:"nombre"`
	ApellidoPat     string `json:"apellidopat"`
	ApellidoMat     string `json:"apellidomat"`
	Genero          string `json:"genero"`
	DNI             string `json:"dni"`
	FechaNacimiento string `json:"fechanacimiento"`
}
