package authorizer

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

func httpGet(url string, headers *map[string]string) (string, error) {
	req, err := http.NewRequest(`GET`, url, nil)
	if err != nil {
		return ``, err
	}
	return doHttpRequest(req, headers)
}

func httpPost(
	url string,
	body string,
	headers *map[string]string,
) (string, error) {
	req, err := http.NewRequest(`POST`, url, strings.NewReader(body))
	if err != nil {
		return ``, err
	}
	return doHttpRequest(req, headers)
}

func doHttpRequest(
	req *http.Request,
	headers *map[string]string,
) (string, error) {
	if headers != nil {
		for header, val := range *headers {
			req.Header.Add(header, val)
		}
	}
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return ``, err
	}
	defer resp.Body.Close()

	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return ``, err
	}
	if resp.StatusCode > 299 {
		return ``, fmt.Errorf(string(respBody))
	}
	return string(respBody), nil
}
