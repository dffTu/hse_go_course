package request

import (
	"encoding/base64"
	"encoding/json"
	"net/http"
	"task2/models"
)

func PrintAPI(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	w.Write([]byte("v1.0.0"))
}

func Decode(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	var encoded models.EncodedString
	err := json.NewDecoder(r.Body).Decode(&encoded)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	data, err := base64.StdEncoding.DecodeString(encoded.Base64)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	json.NewEncoder(w).Encode(models.DecodedString{DecodedFromBase64: string(data)})
}
