package handler

import (
	"net"
	"net/http"
)

func getIP(r *http.Request) string {
	ip := r.Header.Get("x-forwarded-for")
	if ip != "" {
		return ip
	}

	ip = r.Header.Get("x-real-ip")
	if ip != "" {
		return ip
	}

	host, _, err := net.SplitHostPort(r.RemoteAddr)
	if err != nil {
		return r.RemoteAddr
	}

	return host
}

func IPHandler(w http.ResponseWriter, r *http.Request) {
	ip := getIP(r)

	w.Header().Set("Content-Type", "text/plain")
	w.Write([]byte(ip + "\n"))
}
