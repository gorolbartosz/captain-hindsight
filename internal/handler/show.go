package handler

import (
	"fmt"
	"net/http"
)

func Show(w http.ResponseWriter, req *http.Request) {
	w.WriteHeader(200)
	fmt.Fprintln(w, "{}")
}
