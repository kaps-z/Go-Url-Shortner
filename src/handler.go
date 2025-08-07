package main

import (
    "encoding/json"
    "io"
    "net/http"
    "strings"
)

type ShortenRequest struct {
    URL string `json:"url"`
}

func shortenHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var url string

	if r.Header.Get("Content-Type") == "application/x-www-form-urlencoded" {
		url = r.FormValue("url")
	} else {
		body, _ := io.ReadAll(r.Body)
		var req ShortenRequest
		json.Unmarshal(body, &req)
		url = req.URL
	}

	if !strings.HasPrefix(url, "http") {
		if strings.Contains(r.Header.Get("Accept"), "text/html") {
			tmpl.Execute(w, map[string]string{"Error": "Invalid URL format"})
			return
		}
		http.Error(w, "Invalid URL", http.StatusBadRequest)
		return
	}

	shortCode := GenerateShortCode()
	err := SaveURL(shortCode, url)
	if err != nil {
		http.Error(w, "Could not save URL", http.StatusInternalServerError)
		return
	}

	shortURL := "http://localhost:8080/r/" + shortCode

	if strings.Contains(r.Header.Get("Accept"), "text/html") || r.Header.Get("Content-Type") == "application/x-www-form-urlencoded" {
		tmpl.Execute(w, map[string]string{"ShortURL": shortURL})
	} else {
		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte(`{"short_url": "` + shortURL + `"}`))
	}
}

func redirectHandler(w http.ResponseWriter, r *http.Request) {
	code := strings.TrimPrefix(r.URL.Path, "/r/")
	if code == "" {
		http.Error(w, "Not found", http.StatusNotFound)
		return
	}

	original, err := GetOriginalURL(code)
	if err != nil {
		http.Error(w, "URL not found", http.StatusNotFound)
		return
	}

	http.Redirect(w, r, original, http.StatusFound)
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}

	tmpl.Execute(w, nil)
}