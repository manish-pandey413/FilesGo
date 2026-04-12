package config

import (
	"log/slog"
	"net/http"
	"os"
)

type Application struct {
	Logger *slog.Logger
	Server *http.Server
}

func NewApplication(serveMux *http.ServeMux) *Application {

	// logger with time disabled
	lg := slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{
		ReplaceAttr: func(groups []string, a slog.Attr) slog.Attr {
			if a.Key == slog.TimeKey {
				return slog.Attr{}
			}
			return a
		},
	}))

	srvr := &http.Server{
		Addr:    ":8080",
		Handler: serveMux,
	}

	return &Application{
		Logger: lg,
		Server: srvr,
	}
}
