package handler

import (
	"crypto/rand"
	"math/big"
	"net/http"
)

const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789!@#$%^&*()-_=+[]{}|;:,.<>?~"

func generatePassword(length int) (string, error) {
	result := make([]byte, length)

	for i := range result {
		num, err := rand.Int(rand.Reader, big.NewInt(int64(len(charset))))
		if err != nil {
			return "", err
		}
		result[i] = charset[num.Int64()]
	}

	return string(result), nil
}

func Handler(w http.ResponseWriter, r *http.Request) {
	password, err := generatePassword(40)
	if err != nil {
		http.Error(w, "error generating password", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "text/plain")
	w.Write([]byte(password + "\n"))
}
