package gateway

import (
	"fmt"
	"io"
	"io/ioutil"
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
	body, err := ioutil.ReadAll(resp.Body)

	defer resp.Body.Close()

	if err != nil {
		panic(err.Error())
	}

	bodyString := string(body)

	result := StrRemoveAt(bodyString, 0, 1)
	result = Reverse(result)
	result = StrRemoveAt(result, 0, 1)
	result = Reverse(result)

	fmt.Println(result)

	return result
}

func Reverse(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

func StrRemoveAt(s string, index, length int) string {
	return s[:index] + s[index+length:]
}

func Send(payload io.Reader, url string, token string) (string, error) {

	// bearer := "Bearer eyJraWQiOiIxIiwiYWxnIjoiSFMyNTYifQ.eyJ1aWQiOjIsInV0eXBpZCI6MiwiaWF0IjoxNzE3Mjc3ODI4LCJleHAiOjE3MTcyODE0Mjh9.GuG1O8p3kFVL8JAAghGF65mZmSTDA-iKqGuVFQ83Bko"
	bearer := "Bearer " + token

	resp, err := http.NewRequest("POST", url, payload)

	resp.Header.Add("Authorization", bearer)
	resp.Header.Add("Accept", "application/json")

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

	fmt.Println(bodyString)
	fmt.Println("____________________________________________________________________________________________________________________________________________________________")

	return bodyString, nil
}
