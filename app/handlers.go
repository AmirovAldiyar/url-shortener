package app

import (
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"url-shortener/app/models"
)

func (a *App) IndexHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Welcome to Post API")
	}
}

func (a *App) CreateShortUrlHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		req := models.ShortUrlRequest{}
		err := parse(w, r, &req)
		status := http.StatusOK
		if err != nil {
			log.Printf("Cannot parse post body, err=%v \n", err)
			sendResponse(w, r, nil, http.StatusBadRequest)
			return
		}

		short, err := a.DB.GetShortUrl(req.Long)
		if err != nil {
			sendResponse(w, r, nil, http.StatusInternalServerError)
		}

		shortUrl := &models.ShortUrl{
			ID:    0,
			Long:  req.Long,
			Short: short,
		}

		if short == "" {
			short = randomString()
			shortUrl.Short = short
			err = a.DB.CreateShortUrl(shortUrl)
			maxTries := 10
			for err != nil && maxTries > 0 {
				short = randomString()
				shortUrl.Short = short
				err = a.DB.CreateShortUrl(shortUrl)
				maxTries--
			}
			status = http.StatusCreated
		}
		if err != nil {
			log.Printf("Cannot save post in DB. err=%v\n", err)
			sendResponse(w, r, nil, http.StatusInternalServerError)
			return
		}

		resp := mapShortUrlToJSON(shortUrl)
		sendResponse(w, r, resp, status)
	}
}

func (a *App) GetShortUrlHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		long, err := a.DB.GetLongUrl(vars["shortUrl"])
		if err != nil {
			log.Printf("Cannot get long url. err=%v\n", err)
			sendResponse(w, r, nil, http.StatusInternalServerError)
			return
		}

		if long == "" {
			sendResponse(w, r, struct {
				Error string `json:"error"`
			}{Error: "page not found"}, http.StatusNotFound)
			return
		}

		sendResponse(w, r, struct {
			Long string `json:"long"`
		}{Long: long}, http.StatusOK)
	}
}
