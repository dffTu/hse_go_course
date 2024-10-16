package main

import (
	"bytes"
	"context"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"task2/models"
	"time"
)

func getSemanticVersion() {
	request, _ := http.NewRequest(http.MethodGet, "http://localhost:8080/version", nil)
	response, _ := http.DefaultClient.Do(request)
	defer response.Body.Close()
	result, _ := io.ReadAll(response.Body)
	fmt.Println(string(result))
	if string(result) != "v1.0.0" {
		panic("WRONG SEMANTIC VERSION")
	}
}

func decodeString(value string) {
	encoded := models.EncodedString{Base64: base64.StdEncoding.EncodeToString([]byte(value))}
	json_bytes, _ := json.Marshal(encoded)
	request, _ := http.NewRequest(http.MethodPost, "http://localhost:8080/decode", bytes.NewBuffer(json_bytes))
	response, _ := http.DefaultClient.Do(request)
	defer response.Body.Close()
	result, _ := io.ReadAll(response.Body)
	fmt.Println(string(result))
	var decoded models.DecodedString
	json.Unmarshal(result, &decoded)
	if value != decoded.DecodedFromBase64 {
		panic("WRONG DECODING")
	}
}

func hardOperation() {
	request, _ := http.NewRequest(http.MethodGet, "http://localhost:8080/hard-op", nil)
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(time.Second*15))
	defer cancel()
	request = request.WithContext(ctx)
	response, _ := http.DefaultClient.Do(request)
	defer response.Body.Close()
	select {
	case <-ctx.Done():
		fmt.Println(false)
		return
	default:
		fmt.Println("true,", response.StatusCode)
	}
}

func main() {
	getSemanticVersion()
	decodeString("Some message")
	hardOperation()
}
