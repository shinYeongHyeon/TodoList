package api

import (
	"encoding/json"
	"github.com/labstack/gommon/log"
	"io"
	"net/http"
)

func must(err error) {
	if err != nil {
		log.Error("Internal error : ", err)
		panic(internalError)
	}
}

func writeJSON(w http.ResponseWriter, v interface{}) {
	w.Header().Set("Content-Type", "application/json")
	must(json.NewEncoder(w).Encode(v))
}

func parseJSON(r io.Reader, v interface{}) {
	if err := json.NewDecoder(r).Decode(v); err != nil {
		log.Error("Error In parsing json body : ", err)
		panic(malformedInputError)
	}
}