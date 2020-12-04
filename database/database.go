package database

import (
	"database/sql"
	"fmt"
)

/*InitDB permite establecer la conexion a la BD*/
func InitDB() *sql.DB {
	connectionString := "ukf211ji4tzzep6y:xj2mcxWTi5IDXjObwdqc@tcp(balfbddh1nhytmyjvncr-mysql.services.clever-cloud.com:3306)/balfbddh1nhytmyjvncr"
	databaseConnection, err := sql.Open("mysql", connectionString)
	if err != nil {
		fmt.Println("Conexion invalida a la BD")
		panic(err.Error()) // Error Handling = manejo de errores
	}
	return databaseConnection
}
