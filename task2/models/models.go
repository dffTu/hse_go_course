package models

type EncodedString struct {
	Base64 string `json:"inputString"`
}

type DecodedString struct {
	DecodedFromBase64 string `json:"outputString"`
}
