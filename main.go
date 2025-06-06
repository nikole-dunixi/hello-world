package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"time"
)

func handler(w http.ResponseWriter, r *http.Request) {
	hostname, err := os.Hostname()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		io.WriteString(w, err.Error())
		return
	}
	remoteAddr := r.RemoteAddr
	if xForwardedFor, ok := r.Header[`X-Forwarded-For`]; ok && len(xForwardedFor) > 0 {
		remoteAddr = xForwardedFor[0]
	} else if xForwardedHost, ok := r.Header[`X-Forwarded-Host`]; ok && len(xForwardedHost) > 0 {
		remoteAddr = xForwardedHost[0]
	}
	epoch := time.Now().Format(time.RFC3339Nano)
	msg := fmt.Sprintf("Hello, %s! I'm %s! [%s]", remoteAddr, hostname, epoch)
	log.Println(msg)
	io.WriteString(w, msg)
}

func main() {
	http.HandleFunc("/", handler)
	log.Println("listening")
	if err := http.ListenAndServe("0.0.0.0:8080", nil); err != nil {
		log.Fatalf("fatal error: %v", err)
	}
}
