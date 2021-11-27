package database

import (
	"errors"
	"url-shortener/app/models"
)

func (d *DB) CreateShortUrl(p *models.ShortUrl) error {
	res, err := d.db.Exec(insertShortUrlSchema, p.Short, p.Long)
	if err != nil {
		return err
	}

	res.LastInsertId()
	return err
}

func (d *DB) GetShortUrl(longUrl string) (string, error) {
	var shortUrl []string
	err := d.db.Select(&shortUrl, getShortUrlSchema, longUrl)
	if err != nil {
		return "", err
	}

	if len(shortUrl) == 0 {
		return "", nil
	}

	return shortUrl[0], nil
}

func (d *DB) GetLongUrl(shortUrl string) (string, error) {
	var longUrl []string
	err := d.db.Select(&longUrl, getLongUrlSchema, shortUrl)
	if err != nil {
		return "", err
	}

	if len(longUrl) == 0 {
		return "", nil
	}

	return longUrl[0], nil
}

func (d *MemoryDB) CreateShortUrl(p *models.ShortUrl) error {
	if _, exists := d.db[p.Long]; exists {
		return errors.New("already exists, not unique")
	}
	d.db[p.Long] = p.Short
	return nil
}

func (d *MemoryDB) GetShortUrl(longUrl string) (string, error) {
	short, exists := d.db[longUrl]

	if !exists {
		return "", nil
	}

	return short, nil
}

func (d *MemoryDB) GetLongUrl(shortUrl string) (string, error) {
	result := ""
	for key, element := range d.db {
		if element == shortUrl {
			result = key
			break
		}
	}
	return result, nil
}
