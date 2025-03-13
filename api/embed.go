package api

import (
	"embed"
	"github.com/betterde/ects/internal/journal"
	"io/fs"
	"net/http"
)

//go:embed docs/*
var FS embed.FS

func Serve() http.FileSystem {
	dist, err := fs.Sub(FS, "api")
	if err != nil {
		journal.Logger.Panicw("Error mounting front-end static resources!", err)
	}

	return http.FS(dist)
}
