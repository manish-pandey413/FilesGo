package handler

import (
	"io"
	"net/http"
	"os"
	"path/filepath"
)

func ExtractFile(w http.ResponseWriter, r *http.Request) {
	err := r.ParseMultipartForm(100 << 20)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	srcFile, fileHeader, err := r.FormFile("file")
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	defer srcFile.Close()

	fileName := filepath.Base(fileHeader.Filename)

	homePath, err := os.UserHomeDir()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	dst, err := os.Create(homePath + "/FilesGo/" + fileName)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer dst.Close()

	_, err = io.Copy(dst, srcFile)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
