package util

import (
	"errors"
	"io/ioutil"
	"net/http"
	"time"
)

func HttpGet(url string, timeout time.Duration) ([]byte, error) {
	c := http.Client{Timeout: timeout}
	resp, err := c.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return nil, errors.New("status code fail")
	}
	return ioutil.ReadAll(resp.Body)
}
