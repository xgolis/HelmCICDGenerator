package app

import (
	"fmt"
	"net/http"
	"os"
	"os/exec"
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
	fmt.Println(app)
	_, err = os.ReadDir("./" + app.Name)
	if err == nil {
		os.RemoveAll("./" + app.Name)
		fmt.Println("FOUND DIR REMOVING")
	}

	err = os.Mkdir(app.Name, os.ModePerm)
	if err != nil {
		fmt.Printf("error while creating a new dir: %v\n", err)
		return
	}

	err = os.Mkdir(app.Name+"/helm", os.ModePerm)
	if err != nil {
		fmt.Printf("error while creating a new dir: %v\n", err)
		// fmt.Println("error while creating a new dir")
		return
	}

	err = os.Mkdir(app.Name+"/.github", os.ModePerm)
	if err != nil {
		fmt.Printf("error while creating a new dir: %v\n", err)
		// fmt.Println("error while creating a new dir")
		return
	}

	err = createHelmChart(*app)
	if err != nil {
		// fmt.Println("error while creating a new helm chart")
		fmt.Printf("error while creating a new helm chart: %v\n", err)
		return
	}

	err = createCICDPipelines(*app)
	if err != nil {
		fmt.Printf("error while creating a new pipelines: %v\n", err)
		// fmt.Println("error while creating a new pipelines")
		return
	}

	os.MkdirAll("tmp/", 0755)
	// tar := new(archivex.TarFile)
	// tar.Create("tmp/" + app.Name + ".tar")
	// tar.AddAll(app.Name, false)
	// defer tar.Close()
	cmd := exec.Command("tar", "-cf", "tmp/"+app.Name+".tar", "app.Name")

	// Set the working directory for the command (optional)
	// cmd.Dir = "./"

	// Set the standard output and error for the command
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	// Run the command and wait for it to finish
	if err := cmd.Run(); err != nil {
		fmt.Printf("Error: %s\n", err)
		return
	}

	w.Header().Set("Content-Type", "application/x-tar")
	w.Header().Set("Content-Disposition", "attachment; filename=tmp/"+app.Name+".tar")
	http.ServeFile(w, r, "tmp/"+app.Name+".tar")

	os.RemoveAll("./" + app.Name)
	os.RemoveAll("./tmp")
}
