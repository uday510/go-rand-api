package handler

import (
	"crypto/rand"
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
	reqID := FromRequest(r)
	w.Header().Set("X-Request-Id", reqID)

	w.Header().Set("Content-Type", "text/plain; charset=UTF-8")
	w.Header().Set("Cache-Control", "no-store, no-cache, must-revalidate, proxy-revalidate")
	w.Header().Set("Pragma", "no-cache")
	w.Header().Set("Expires", "0")
	w.Header().Set("X-Content-Type-Options", "nosniff")

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

	w.WriteHeader(http.StatusOK)
	_, _ = w.Write([]byte(password + "\n"))
}
