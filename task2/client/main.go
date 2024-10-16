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

const (
	server_addr      string        = "http://localhost:8080"
	semantic_version string        = "v1.0.0"
	hard_op_timeout  time.Duration = time.Duration(time.Second * 15)
)

func getSemanticVersion() error {
	request, err := http.NewRequest(http.MethodGet, server_addr+"/version", nil)
	if err != nil {
		return err
	}
	response, err := http.DefaultClient.Do(request)
	if err != nil {
		return err
	}
	defer response.Body.Close()
	result, err := io.ReadAll(response.Body)
	if err != nil {
		return err
	}
	fmt.Println(string(result))
	if string(result) != semantic_version {
		return fmt.Errorf("WRONG SEMANTIC VERSION")
	}
	return nil
}

func decodeString(value string) error {
	encoded := models.EncodedString{Base64: base64.StdEncoding.EncodeToString([]byte(value))}
	json_bytes, err := json.Marshal(encoded)
	if err != nil {
		return err
	}
	request, err := http.NewRequest(http.MethodPost, server_addr+"/decode", bytes.NewBuffer(json_bytes))
	if err != nil {
		return err
	}
	response, err := http.DefaultClient.Do(request)
	if err != nil {
		return err
	}
	defer response.Body.Close()
	result, err := io.ReadAll(response.Body)
	if err != nil {
		return err
	}
	fmt.Println(string(result))
	var decoded models.DecodedString
	json.Unmarshal(result, &decoded)
	if value != decoded.DecodedFromBase64 {
		return fmt.Errorf("WRONG DECODING")
	}
	return nil
}

func hardOperation() error {
	request, err := http.NewRequest(http.MethodGet, server_addr+"/hard-op", nil)
	if err != nil {
		return err
	}
	ctx, cancel := context.WithTimeout(context.Background(), hard_op_timeout)
	defer cancel()
	request = request.WithContext(ctx)
	response, err := http.DefaultClient.Do(request)
	if err != nil {
		return err
	}
	defer response.Body.Close()
	select {
	case <-ctx.Done():
		fmt.Println(false)
		return nil
	default:
		fmt.Println("true,", response.StatusCode)
		return nil
	}
}

func main() {
	err := getSemanticVersion()
	if err != nil {
		fmt.Println(err.Error())
	}
	err = decodeString("Some message")
	if err != nil {
		fmt.Println(err.Error())
	}
	err = hardOperation()
	if err != nil {
		fmt.Println(err.Error())
	}
}
