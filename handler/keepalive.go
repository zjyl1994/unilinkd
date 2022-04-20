package handler

import (
	"io/ioutil"
	"path/filepath"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/patrickmn/go-cache"
	"github.com/zjyl1994/unilinkd/util"
)

var cacheInstance = cache.New(5*time.Minute, 10*time.Minute)

func KeepAliveHandler(c *fiber.Ctx, url string) error {
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
		return err
	}
	return c.Send(data)
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
