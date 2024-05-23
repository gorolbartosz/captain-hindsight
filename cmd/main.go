package main

import (
	_ "embed"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/handlers"
	"github.com/gorolbartosz/captain-hindsight/internal/app"
	"github.com/gorolbartosz/captain-hindsight/internal/handler"
	"github.com/gorolbartosz/captain-hindsight/internal/image"

	"github.com/mattermost/mattermost-plugin-apps/apps"
	"github.com/mattermost/mattermost-plugin-apps/utils/httputils"
)

func main() {
	http.HandleFunc("/favicon.ico", httputils.DoHandleData("image/png", image.Logo))
	http.HandleFunc("/manifest.json", httputils.DoHandleJSON(app.Manifest))
	http.HandleFunc("/bindings", httputils.DoHandleJSON(apps.NewDataResponse(app.Bindings)))
	http.HandleFunc("/ping", ping)

	http.HandleFunc("/static/logo.png", httputils.DoHandleData("image/png", image.Logo))

	http.HandleFunc("/list", handler.List)
	http.HandleFunc("/show", handler.Show)
	http.HandleFunc("/trigger", handler.Trigger)

	addr := "127.0.0.1:8066"
	fmt.Println("Captain we need your help at", addr)
	log.Fatal(http.ListenAndServe(addr, handlers.LoggingHandler(os.Stdout, http.DefaultServeMux)))
}

func ping(w http.ResponseWriter, req *http.Request) {
	w.WriteHeader(200)
	fmt.Fprintln(w, "{}")
}
