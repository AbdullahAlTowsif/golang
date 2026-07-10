package util

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
)

type Header struct {
	Alg string `json:"alg"`
	Typ string `json:"typ"`
}

type Payload struct {
	Sub         int `json:"sub"` // user id
	FirstName   string `json:"first_name"`
	LastName    string `json:"last_name"`
	Email       string `json:"email"`
	IsShopOwner bool   `json:"is_shop_owner"`
}

func CreateJwt(secret string, data Payload) (string, error) {
	header := Header{
		Alg: "HS256",
		Typ: "JWT",
	}

	byteArrHeader, err := json.Marshal(header) // converted header to byte
	if err != nil {
		return "", err
	}
	headerB64 := base64UrlEncode(byteArrHeader) // converting header byte to base64

	byteArrData, err := json.Marshal(data) // converted data/payload to byte
	if err != nil {
		return "", err
	}
	payloadB64 := base64UrlEncode(byteArrData) // converting payload byte to base64

	message := headerB64 + "." + payloadB64

	// now signature
	byteArrSecret := []byte(secret)
	byteArrMessage := []byte(message)

	// made a secret key
	h := hmac.New(sha256.New, byteArrSecret)
	h.Write(byteArrMessage)

	signature := h.Sum(nil) // got the hash of the message using the secret key
	signatureB64 := base64UrlEncode(signature) // now converting the secret key to base64

	jwt := headerB64 + "." + payloadB64 + "." + signatureB64
	return jwt, nil
}


func base64UrlEncode(data []byte) string {
	return base64.URLEncoding.WithPadding(base64.NoPadding).EncodeToString(data)
}
