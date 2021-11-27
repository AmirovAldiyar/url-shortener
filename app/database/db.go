package database

import (
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"log"
	"url-shortener/app/models"
)

type UrlShortDB interface {
	Open() error
	Close() error
	CreateShortUrl(p *models.ShortUrl) error
	GetShortUrl(longUrl string) (string, error)
	GetLongUrl(shortUrl string) (string, error)
}

type MemoryDB struct {
	db map[string]string
}

type DB struct {
	db *sqlx.DB
}

func (d *DB) Open() error {
	initConfig()
	pg, err := sqlx.Open("postgres", pgConnStr)
	if err != nil {
		return err
	}
	log.Println("Connected to Database!")

	pg.MustExec(createSchema)

	d.db = pg

	return nil
}

func (d *DB) Close() error {
	return d.db.Close()
}

func (d *MemoryDB) Open() error {
	d.db = make(map[string]string)
	log.Println("Connected to Database!")
	return nil
}

func (d *MemoryDB) Close() error {
	return nil
}
