package webhandlers

import "net/http"

// HelloWorldHandler simply returns an alive signal.
func HelloWorldHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`{"status": "alive"}`))
}
