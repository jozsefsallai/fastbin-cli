package utils

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/jozsefsallai/fastbin-cli/config"
)

// Upload will upload a file to the remote server.
func Upload(input string) (string, error) {
	conf := config.GetConfig()
	fmt.Println(input)

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

	return string(body), nil
}
