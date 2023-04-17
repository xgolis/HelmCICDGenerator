package app

import (
	"fmt"
	"net/http"
)

func MakeHandlers() *http.ServeMux {
	mux := *http.NewServeMux()
	mux.HandleFunc("/", getCharts)
	return &mux
	// a.Server.HandleFunc("/", getGit(w http.ResponseWriter, req *http.Request)})
}

func getCharts(w http.ResponseWriter, r *http.Request) {
	app, err := getUserApp(r)
	if err != nil {
		fmt.Fprintf(w, "error while parsing json from request body: %v", err)
		return
	}

	helmGenerator.createHelmChart(app)
}
