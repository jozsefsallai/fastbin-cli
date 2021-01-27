package utils

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/jozsefsallai/fastbin-cli/config"
)

// RequestResponse is the structure of a potential JSON response
// coming from the server
type RequestResponse struct {
	Ok     bool   `json:"ok"`
	Error  string `json:"error"`
	Key    string `json:"key"`
	Secret string `json:"secret"`
}

// Upload will upload a file to the remote server.
func Upload(input string) (string, string, error) {
	conf := config.GetConfig()

	url := conf.Server + "/documents"
	payload := bytes.NewBuffer([]byte(input))
	req, err := http.NewRequest("POST", url, payload)
	req.Header.Set("Content-Type", "text/plain")

	if len(conf.Key) > 0 {
		req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", conf.Key))
	}

	if err != nil {
		return "", "", err
	}

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return "", "", err
	}

	defer res.Body.Close()

	body, _ := ioutil.ReadAll(res.Body)

	var jsonResponse RequestResponse
	json.Unmarshal(body, &jsonResponse)

	if jsonResponse.Ok == false {
		if len(jsonResponse.Error) > 0 {
			return "", "", errors.New(jsonResponse.Error)
		}

		return "", "", errors.New("failed to upload the snippet")
	}

	return string(jsonResponse.Key), string(jsonResponse.Secret), nil
}
