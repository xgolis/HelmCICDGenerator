package main

import (
	"fmt"
	"net/http"

	"github.com/xgolis/HelmCICDGenerator/cmd/HelmCICDGenerator/app"
)

func main() {
	app := app.NewApp()

	fmt.Printf("[Server] Up and running on %s\n", app.Server.Addr)
	http.ListenAndServe(app.Server.Addr, app.Server.Handler)
}
