package db

import (
	"fmt"
	_ "github.com/joho/godotenv/autoload"
	"os"
)

var (
	host     = getenv("HOST", "localhost")
	port     = getenv("PORT", "")
	user     = getenv("USER", "postgres")
	password = getenv("PASSWORD", "")
	dbName   = getenv("NAME_BD", "")
	PortHttp = getenv("PORT_HTTP", "")
)

var ConnStr = fmt.Sprintf("postgres://%s:%s@%s/%s?sslmode=disable", user, password, host, dbName)

func getenv(key, defaultValue string) string {
	value, defined := os.LookupEnv(key)
	if !defined {
		return defaultValue
	}
	return value
}
