package handler

import (
	"fmt"
	"net/http"
)

func Trigger(w http.ResponseWriter, req *http.Request) {
	w.WriteHeader(200)
	fmt.Fprintln(w, "{}")
}
