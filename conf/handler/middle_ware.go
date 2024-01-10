package handler

import (
	"bytes"
	"github.com/sirupsen/logrus"
	"io"
	"net/http"
)

var excludeApiLog = []string{"/actuator"}

func HandleMiddleWare(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if contains(excludeApiLog, r.URL.Path) {
			next.ServeHTTP(w, r)
			return
		}
		body, err := io.ReadAll(r.Body)
		if err != nil {
			http.Error(w, "Error reading request body", http.StatusBadRequest)
			return
		}
		logrus.Info("---[Inbound] Start ", r.Method, r.URL.Path, " ---")
		logrus.Info("[Inbound] Request body: \n", string(body))

		r.Body = io.NopCloser(bytes.NewBuffer(body))

		// Call the next handler in the chain
		next.ServeHTTP(w, r)

		logrus.Info("---[Inbound] Finish ", r.Method, r.URL.Path, " ---")
	})
}
func contains(slice []string, item string) bool {
	for _, i := range slice {
		if i == item {
			return true
		}
	}
	return false
}
