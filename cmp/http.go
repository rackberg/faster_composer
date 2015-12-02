package cmp

import (
	"net/http"
	"io/ioutil"
)

func GetHttpResponseBody (url string) (content []byte, err error) {
	response, err := http.Get(url)
	defer response.Body.Close()
	content, err = ioutil.ReadAll(response.Body)

	return content, err
}