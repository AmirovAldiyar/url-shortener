package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"
	"testing"
	"time"
	"url-shortener/app"
	"url-shortener/app/database"
	"url-shortener/app/models"
)

func TestHandler(t *testing.T) {
	tt := []struct {
		name       string
		query      string
		want       string
		method     string
		firstUrl   string
		secondUrl  string
		statusCode int
	}{
		{
			name:       "not exists",
			query:      "+123123123",
			want:       "{\"error\":\"page not found\"}\n",
			method:     http.MethodGet,
			statusCode: http.StatusNotFound,
		}, {
			name:       "exists",
			query:      "",
			want:       "{\"long\":\"www.youtube.com\"}\n",
			method:     http.MethodGet,
			statusCode: http.StatusOK,
		}, {
			name:       "same urls",
			firstUrl:   "www.youtube.com",
			secondUrl:  "www.youtube.com",
			method:     http.MethodPost,
			statusCode: http.StatusOK,
			want:       "equal",
		}, {
			name:       "different urls",
			firstUrl:   "www.youtube.com",
			secondUrl:  "www.google.com",
			method:     http.MethodPost,
			statusCode: http.StatusCreated,
			want:       "not equal",
		},
	}

	app := app.New()
	app.DB = &database.MemoryDB{}
	err := app.DB.Open()
	check(err)

	defer app.DB.Close()

	http.HandleFunc("/", app.Router.ServeHTTP)

	log.Println("Start serving...")
	go http.ListenAndServe(":9000", nil)
	time.Sleep(2 * time.Second)
	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			switch tc.method {
			case http.MethodGet:
				body := `{"long": "www.youtube.com"}`
				query := tc.query
				client := http.Client{}

				resp, err := client.Post("http://localhost:9000/api/shorten", "application/json", strings.NewReader(body))
				buf := new(bytes.Buffer)
				buf.ReadFrom(resp.Body)

				if query == "" {
					unmarshalled := models.ShortUrlRequest{}
					err = json.Unmarshal([]byte(buf.String()), &unmarshalled)

					if err != nil {
						t.Errorf("%s", err)
					}
					query = unmarshalled.Short
				}

				resp, err = client.Get("http://localhost:9000/" + query)

				buf = new(bytes.Buffer)
				buf.ReadFrom(resp.Body)
				respBody := models.JsonShortUrl{}
				json.Unmarshal([]byte(buf.String()), &respBody)

				if resp.StatusCode != tc.statusCode {
					t.Errorf("Want status '%d', got '%d'", tc.statusCode, resp.StatusCode)
				}
				if buf.String() != tc.want {
					t.Errorf("Want '%s', got '%s'", tc.want, buf.String())
				}
			case http.MethodPost:
				body := fmt.Sprintf(`{"long": "%s"}`, tc.firstUrl)
				client := http.Client{}

				resp, err := client.Post("http://localhost:9000/api/shorten", "application/json", strings.NewReader(body))
				buf := new(bytes.Buffer)
				buf.ReadFrom(resp.Body)
				unmarshalled := models.ShortUrlRequest{}
				err = json.Unmarshal([]byte(buf.String()), &unmarshalled)

				if err != nil {
					t.Errorf("%s", err)
				}

				firstResult := unmarshalled.Short

				body = fmt.Sprintf(`{"long": "%s"}`, tc.secondUrl)
				resp, err = client.Post("http://localhost:9000/api/shorten", "application/json", strings.NewReader(body))
				buf = new(bytes.Buffer)
				buf.ReadFrom(resp.Body)
				unmarshalled = models.ShortUrlRequest{}
				err = json.Unmarshal([]byte(buf.String()), &unmarshalled)

				if err != nil {
					t.Errorf("%s", err)
				}

				secondResult := unmarshalled.Short

				if tc.want == "equal" && firstResult != secondResult {
					t.Errorf("Short Urls should be equal, got (%s, %s)", firstResult, secondResult)
				}
				if tc.want == "not equal" && firstResult == secondResult {
					t.Errorf("Short Urls should not be equal, got (%s, %s)", firstResult, secondResult)
				}
				if tc.statusCode != resp.StatusCode {
					t.Errorf("Status code should be %v, got %v", tc.statusCode, resp.StatusCode)
				}
			}

		})
	}
}
