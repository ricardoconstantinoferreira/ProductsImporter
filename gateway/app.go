package gateway

import (
	"import/utils"
	"io"
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

	defer resp.Body.Close()

	if err != nil {
		panic(err.Error())
	}

	bodyString := string(body)

	result := utils.StrRemoveAt(bodyString, 0, 1)
	result = utils.Reverse(result)
	result = utils.StrRemoveAt(result, 0, 1)
	result = utils.Reverse(result)

	return result
}

func Send(payload io.Reader, url string, token string) string {

	bearer := "Bearer " + token

	client := &http.Client{}
	req, err := http.NewRequest("POST", url, payload)

	if err != nil {
		panic(err.Error())
	}

	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", bearer)

	resp, err := client.Do(req)
	if err != nil {
		panic(err.Error())
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		panic(err.Error())
	}

	return string(body)
}
