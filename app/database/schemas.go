package database

const createSchema = `
CREATE TABLE IF NOT EXISTS short_url
(
	id SERIAL PRIMARY KEY,
	short TEXT UNIQUE,
	long TEXT UNIQUE
)
`

var insertShortUrlSchema = `
INSERT INTO short_url(short, long) VALUES($1, $2) RETURNING id
`

var getShortUrlSchema = `
SELECT short FROM short_url WHERE long=$1
`

var getLongUrlSchema = `
SELECT long FROM short_url WHERE short=$1
`
