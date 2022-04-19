package server

import (
	"fmt"
	"net/http"
	"time"

	"github.com/zjyl1994/unilinkd/config"
)

type UnilinkdServer struct{}

func (UnilinkdServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		w.WriteHeader(http.StatusMethodNotAllowed)
		writePage(w, "Not Allowed", fmt.Sprintf("Method '%s' not allow on server.", r.Method))
		return
	}
	if r.URL.Path == "" || r.URL.Path == "/" {
		w.Write(indexData)
		return
	}
	link, found := config.Links[r.URL.Path]
	if !found {
		w.WriteHeader(http.StatusNotFound)
		writePage(w, "Not Found", fmt.Sprintf("Path '%s' not found on server.", r.URL.Path))
		return
	}
	if !link.Expire.IsZero() && time.Now().After(link.Expire) {
		w.WriteHeader(http.StatusNotFound)
		writePage(w, "Not Found", fmt.Sprintf("Path '%s' not found on server.", r.URL.Path))
		return
	}
	proc, found := modeProc[link.Mode]
	if !found {
		w.WriteHeader(http.StatusNotImplemented)
		writePage(w, "Not Implemented", fmt.Sprintf("Link mode '%s' not implemented.", link.Mode))
		return
	}
	proc(w, r, link.Url)
}
