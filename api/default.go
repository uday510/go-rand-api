package handler

import (
	"apis/lib"
	"net/http"
)

func Handler(w http.ResponseWriter, r *http.Request) {
	lib.WriteTextResponse(w, r, "welcome to earth, human. you seem lost.")
}
