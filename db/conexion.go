package db

import (
	"database/sql"
	_ "github.com/lib/pq"
)

var Db *sql.DB

func AbrirConexion() {
	conexion, err := sql.Open("postgres", ConnStr)
	if err != nil {
		panic(err)
	}
	Db = conexion
}
func CerrarConexion() {
	Db.Close()
}
