package handler

import "net/http"

func FileHandler(w http.ResponseWriter, r *http.Request, url string) {
	http.ServeFile(w, r, url)
}

func UrlHandler(w http.ResponseWriter, r *http.Request, url string) {
	http.Redirect(w, r, url, http.StatusFound)
}
