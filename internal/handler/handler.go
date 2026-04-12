package handler

import (
	"io"
	"net/http"
	"os"
	"path/filepath"

	"github.com/manish-pandey413/FilesGo/internal/config"
)

func ExtractFile(w http.ResponseWriter, r *http.Request) {
	app := &config.Application{}

	err := r.ParseMultipartForm(100 << 20)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		app.Logger.Error(err.Error(), "method", r.Method, "uri", r.URL.RequestURI())
		return
	}

	srcFile, fileHeader, err := r.FormFile("file")
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		app.Logger.Error(err.Error(), "method", r.Method, "uri", r.URL.RequestURI())
		return
	}
	defer srcFile.Close()

	fileName := filepath.Base(fileHeader.Filename)

	homePath, err := os.UserHomeDir()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		app.Logger.Error(err.Error(), "method", r.Method, "uri", r.URL.RequestURI())
		return
	}

	dst, err := os.Create(homePath + "/FilesGo/" + fileName)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		app.Logger.Error(err.Error(), "method", r.Method, "uri", r.URL.RequestURI())
		return
	}
	defer dst.Close()

	_, err = io.Copy(dst, srcFile)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		app.Logger.Error(err.Error(), "method", r.Method, "uri", r.URL.RequestURI())
		return
	}
}
