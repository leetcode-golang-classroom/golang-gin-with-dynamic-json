package util

import (
	"encoding/json"
	"log"
	"net/http"
)

func FailOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
	}
}

func CloseChannel(ch chan error) {
	if _, ok := <-ch; ok {
		close(ch)
	}
}

func WriteJSON(w http.ResponseWriter, status int, value any) error {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(status)
	return json.NewEncoder(w).Encode(value)
}

func WriteError(w http.ResponseWriter, status int, err error) {
	errResp := WriteJSON(w, status, map[string]string{"error": err.Error()})
	if errResp != nil {
		log.Fatal(errResp)
	}
}
