package handler

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"
	"time"

	"github.com/patrickmn/go-cache"
	"github.com/zjyl1994/unilinkd/util"
)

var cacheInstance = cache.New(5*time.Minute, 10*time.Minute)

func KeepAliveHandler(w http.ResponseWriter, r *http.Request, url string) {
	data, err := innerCacheGet(url, func() ([]byte, error) {
		cacheFilename := filepath.Join("data", util.MD5String(url))
		data, err := util.HttpGet(url, 3*time.Second)
		if err != nil {
			return ioutil.ReadFile(cacheFilename)
		} else {
			return data, ioutil.WriteFile(cacheFilename, data, 0644)
		}
	})
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintln(os.Stderr, err.Error())
		return
	}
	w.Write(data)
}

func innerCacheGet(key string, fn func() ([]byte, error)) ([]byte, error) {
	if data, found := cacheInstance.Get(key); found {
		return data.([]byte), nil
	}
	data, err := fn()
	if err != nil {
		return nil, err
	}
	cacheInstance.Set(key, data, cache.DefaultExpiration)
	return data, nil
}
