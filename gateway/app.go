package gateway

import (
	"io"
	"log"
	"net/http"
)

func TokenAdmin(payload io.Reader, url string) string {

	resp, err := http.Post(url,
		"application/json", payload)

	if err != nil {
		panic(err.Error())
	}

	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)

	if err != nil {
		panic(err.Error())
	}

	bodyString := string(body)

	return bodyString
}

func Send(payload io.Reader, url string, token string) (string, error) {

	resp, err := http.Post(url,
		"application/json", payload)

	bearer := "Bearer " + token
	resp.Header.Add("Authentication", bearer)

	if err != nil {
		log.Printf("Request Failed: %s", err)
		return "", err
	}

	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)

	if err != nil {
		log.Printf("Read Failed: %s", err)
		return "", err
	}

	bodyString := string(body)

	return bodyString, nil
}
