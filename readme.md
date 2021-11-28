# Url Shortener

## Routes

- `/api/shorten` - POST method - Body should contain 'long' field which contain url that will be shorten
    - Request example: `http://localhost:9000/api/shorten`
        - Body: `{"long": "www.123123.com"}`
    - Response example: `{
      "id": 0,
      "long": "www.123123.com",
      "short": "qS_PiZYmLF"
      }`

- `/{shortUrl}` - GET method - Responding the long version of the url as a json with the field 'long'
    - Request example: `http://localhost:9000/qS_PiZYmLF`
    - Response example: `{
      "long": "www.123123.com"
      }`

## Flags

- `db=$dbType` - decides what will be used as a store: postgres/in-memory. Allowed values are 'postgres' or 'memory'.
  Ex: db=memory
- Ex: `go run main.go --db=memory`

## Docker

- You can pull docker-image from [here](https://hub.docker.com/r/aldeeyar/docker-url-shortener)
- Specify variable 'DB' if you want to use in-memory storage
- Ex: ```docker run -d -t -i -e POSTGRES_PASSWORD=your_password -e DB=memory -p 9000:9000 docker-url-shortener```

## Environment Variables

- ```POSTGRES_USERNAME``` - database user's username (default: 'postgres')
- ```POSTGRES_PASSWORD``` - database user's password (default: empty string)
- ```POSTGRES_HOST``` - database host (default: 'localhost')
- `POSTGRES_TABLE` - database name (default: 'postgres')
- `POSTGRES_PORT` - database port (default: '5432')

## Credits

Telegram: @AmirovAldiyar

E-mail: amirovaldiar@gmail.com