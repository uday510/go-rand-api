package handler

import (
	"crypto/rand"
	"go-rand-api/internal"
	"math/big"
	"net/http"
	"strconv"
)

const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789!@#$%^&*()-_=+[]{}|;:,.<>?~"

func generatePassword(length int) (string, error) {
	result := make([]byte, length)

	mx := big.NewInt(int64(len(charset)))

	for i := range result {
		num, err := rand.Int(rand.Reader, mx)
		if err != nil {
			return "", err
		}
		result[i] = charset[num.Int64()]
	}

	return string(result), nil
}

func RandHandler(w http.ResponseWriter, r *http.Request) {
	length := 40
	if l := r.URL.Query().Get("length"); l != "" {
		if parsed, err := strconv.Atoi(l); err == nil && parsed > 0 && parsed <= 1024 {
			length = parsed
		}
	}

	password, err := generatePassword(length)
	if err != nil {
		http.Error(w, "error generating password", http.StatusInternalServerError)
		return
	}

	internal.WriteTextResponse(w, r, password)
}
