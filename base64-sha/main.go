package main

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"fmt"
)

func main() {
    // base64 encoding
    // string --> (image, file, bytes) ---> Base64 (A-Z,a-z,0-9,+,/)
	var s string
	s = "hello world"

	byteArr := []byte(s)
	fmt.Println(s)
    fmt.Println(byteArr)

    enc := base64.URLEncoding
    enc = enc.WithPadding(base64.NoPadding)
    b64Str := enc.EncodeToString(byteArr)

    fmt.Println(b64Str)

    decodedStr, err := enc.DecodeString(b64Str)
    if err != nil {
        fmt.Println(err)
        return
    }
    fmt.Println(decodedStr)

    fmt.Println("=================")

    // SHA = Secure Hash Algorithm Encoding
    data := []byte("Hello")
    hash := sha256.Sum256(data)

    fmt.Println("Hash after 256:", hash) // byte array

    fmt.Println("=================")

    // HMAC = Hash-based Message Authentication Code
    // text + secret key --> HMAC ---> Hash (Output)
    // HMAC-SHA-256
    secret := []byte("my-secret")
    message := []byte("Hello World")

    h := hmac.New(sha256.New, secret)
    h.Write(message)

    text := h.Sum(nil)
    fmt.Println(text)
}
