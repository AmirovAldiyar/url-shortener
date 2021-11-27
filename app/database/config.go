package database

import (
	"fmt"
	"os"
)

var (
	dbUsername = "postgres"
	dbPassword = ""
	dbHost     = "localhost"
	dbTable    = "postgres"
	dbPort     = "5432"
	pgConnStr  = fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=disable",
		dbHost, dbPort, dbUsername, dbTable, dbPassword)
)

func initConfig() {
	if password, exists := os.LookupEnv("POSTGRES_PASSWORD"); exists {
		dbPassword = password
	}
	if username, exists := os.LookupEnv("POSTGRES_USERNAME"); exists {
		dbUsername = username
	}
	if host, exists := os.LookupEnv("POSTGRES_HOST"); exists {
		dbHost = host
	}
	if table, exists := os.LookupEnv("POSTGRES_TABLE"); exists {
		dbTable = table
	}
	if port, exists := os.LookupEnv("POSTGRES_PORT"); exists {
		dbPort = port
	}
	pgConnStr = fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=disable",
		dbHost, dbPort, dbUsername, dbTable, dbPassword)
}
