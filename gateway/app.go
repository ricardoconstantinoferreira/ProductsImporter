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

	req, err := http.NewRequest("POST", url, payload)

	req.Header.Add("Authorization", bearer)
	req.Header.Add("Accept", "application/json")

	if err != nil {
		panic(err.Error())
	}

	client := &http.Client{}
	response, err := client.Do(req)

	if err != nil {
		panic(err.Error())
	}

	return response.Status
}
