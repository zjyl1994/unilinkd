package server

import (
	"net/http"

	"github.com/zjyl1994/unilinkd/handler"
)

type modeProcHandler func(http.ResponseWriter, *http.Request, string)

var modeProc = map[string]modeProcHandler{
	"file":      handler.FileHandler,
	"url":       handler.UrlHandler,
	"keepalive": handler.KeepAliveHandler,
	"s3":        handler.S3Handler,
}
