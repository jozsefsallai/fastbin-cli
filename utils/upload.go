package utils

import (
	"bytes"
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"

	"github.com/jozsefsallai/fastbin-cli/config"
)

type RequestResponse struct {
	Ok bool `json:"ok"`
	Error string `json:"error"`
	Key string `json:"key"`
}

// Upload will upload a file to the remote server.
func Upload(input string) (string, error) {
	conf := config.GetConfig()

	url := conf.Server + "/documents"
	payload := bytes.NewBuffer([]byte(input))
	req, err := http.NewRequest("POST", url, payload)
	req.Header.Set("Content-Type", "text/plain")
	if err != nil {
		return "", err
	}

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return "", err
	}

	defer res.Body.Close()

	body, _ := ioutil.ReadAll(res.Body)
	if err != nil {
		return "", err
	}

	var jsonResponse RequestResponse
	json.Unmarshal(body, &jsonResponse)

	if jsonResponse.Ok == false {
		if len(jsonResponse.Error) > 0 {
			return "", errors.New(jsonResponse.Error)
		}

		return "", errors.New("Failed to upload the snippet.")
	}

	return string(jsonResponse.Key), nil
}
