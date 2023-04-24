package app

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"time"
)

type App struct {
	Server *http.Server
}

type UserApp struct {
	Name     string `json:"name"`
	UserName string `json:"username"`
	Image    string `json:"image"`
	Port     string `json:"port"`
}

func NewApp() *App {

	mux := MakeHandlers()

	return &App{
		Server: &http.Server{
			Addr:           ":8080",
			Handler:        mux,
			ReadTimeout:    50 * time.Second,
			WriteTimeout:   50 * time.Second,
			MaxHeaderBytes: 1 << 20,
		},
	}
}

func getUserApp(r *http.Request) (*UserApp, error) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return nil, err
	}

	var userApp UserApp
	err = json.Unmarshal(body, &userApp)
	if err != nil {
		return nil, err
	}

	return &userApp, nil
}
