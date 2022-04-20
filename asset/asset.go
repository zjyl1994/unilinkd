package asset

import (
	"embed"
	"net/http"
)

//go:embed *.*
var Assets embed.FS
var HttpAssets http.FileSystem

func init() {
	HttpAssets = http.FS(Assets)
}
