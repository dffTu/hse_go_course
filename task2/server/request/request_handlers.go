package request

import "net/http"

func PrintAPI(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		w.Write([]byte("v1.0.0"))
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}
