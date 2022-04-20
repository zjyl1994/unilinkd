package handler

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/zjyl1994/unilinkd/config"
)

type modeProcHandler func(*fiber.Ctx, string) error

var modeProc = map[string]modeProcHandler{
	"file":      FileHandler,
	"url":       UrlHandler,
	"keepalive": KeepAliveHandler,
	"s3":        S3Handler,
}

func FileHandler(c *fiber.Ctx, url string) error {
	return c.SendFile(url)
}

func UrlHandler(c *fiber.Ctx, url string) error {
	return c.Redirect(url)
}

func CodeHandler(c *fiber.Ctx, code string) error {
	link, found := config.Links[code]
	if !found {
		return fiber.NewError(http.StatusNotFound, fmt.Sprintf("Path '%s' not found on server.", c.OriginalURL()))
	}
	if !link.Expire.IsZero() && time.Now().After(link.Expire) {
		return fiber.NewError(http.StatusNotFound, fmt.Sprintf("Path '%s' not found on server.", c.OriginalURL()))
	}
	proc, found := modeProc[link.Mode]
	if !found {
		return fiber.NewError(http.StatusNotImplemented, fmt.Sprintf("Link mode '%s' not implemented.", link.Mode))
	}
	return proc(c, link.Url)
}
