package request

import (
	"encoding/base64"
	"encoding/json"
	"math/rand/v2"
	"net/http"
	"task2/models"
	"time"
)

const (
	semantic_version string = "v1.0.0"
)

func PrintAPI(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	w.Write([]byte(semantic_version))
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

func HardOperation(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	time.Sleep(time.Duration(rand.IntN(11)+10) * time.Second)
	if rand.IntN(2) == 0 {
		w.WriteHeader(500 + rand.IntN(27))
	} else {
		w.WriteHeader(200)
	}
}
