package handlers

import (
	"errors"
	"net/http"
	"regexp"
	"strings"

	"github.com/MihailSergeenkov/shortener/internal/app/storage"
)

func FetchHandler(w http.ResponseWriter, r *http.Request) {
	re := regexp.MustCompile(`^/\w{8}$`)
	key := re.FindString(r.URL.Path)

	if key == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	u, err := storage.FetchURL(strings.TrimLeft(key, "/"))

	if err != nil {
		if errors.Is(err, storage.ErrURLNotFound) {
			w.WriteHeader(http.StatusNotFound)
			return
		}

		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Location", u)
	w.WriteHeader(http.StatusTemporaryRedirect)
}