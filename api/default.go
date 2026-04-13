package handler

import "net/http"

func Handler(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "text/plain")
	w.Write([]byte("welcome to earth, human. you seem lost.\n"))

}
