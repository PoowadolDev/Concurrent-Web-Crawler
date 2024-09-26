package main

import (
	"io/ioutil"
	"net/http"
)

func fetch(url string) (string, error) {
	resp, err := http.Get(url)

	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	return string(bodyBytes), nil
}
