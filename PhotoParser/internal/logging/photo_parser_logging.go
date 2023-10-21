package logging

import (
	"log"
	"net/http"
	"time"
)

func Logging(postHandler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		log.Printf("Request: %s %s %s\n", r.Method, r.RequestURI, start.Format(time.RFC822))
		defer log.Printf("Response: %s %s %s %s\n", r.Method, r.RequestURI, time.Now().Format(time.RFC822), time.Since(start))
		if r.Method == http.MethodPost {
			postHandler.ServeHTTP(w, r)
			return
		}
		w.WriteHeader(http.StatusMethodNotAllowed)
	})
}
