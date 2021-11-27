package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"url-shortener/app"
	"url-shortener/app/database"
)

func main() {
	dbType := flag.String("db", "postgres", "database type")
	flag.Parse()

	app := app.New()
	if *dbType == "postgres" {
		app.DB = &database.DB{}
	} else if *dbType == "memory" {
		app.DB = &database.MemoryDB{}
	} else {
		fmt.Printf("Expected db type 'postgres' or 'memory', got '%s'\n", *dbType)
		return
	}
	err := app.DB.Open()
	check(err)

	defer app.DB.Close()

	http.HandleFunc("/", app.Router.ServeHTTP)

	log.Println("Start serving...")
	err = http.ListenAndServe(":9000", nil)
	check(err)
}

func check(err error) {
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}
}
