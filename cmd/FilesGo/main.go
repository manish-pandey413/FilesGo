package main

import (
	"flag"
	"net"
	"net/http"
	"os"

	"github.com/manish-pandey413/FilesGo/internal/config"
	"github.com/manish-pandey413/FilesGo/internal/handler"
)

func main() {

	webMode := flag.Bool("webMode", false, "Set true to use web mode.")
	path := flag.String("dir", "./", "Directory to share")
	flag.Parse()

	normalMux := http.NewServeMux()

	var (
		webModeMux *http.ServeMux
		fileServer http.Handler
	)

	if *webMode {
		webModeMux = http.NewServeMux()

		fileServer = http.FileServer(http.Dir(*path))
		webModeMux.Handle("GET /", fileServer)
	}

	serveMux := normalMux
	if *webMode {
		serveMux = webModeMux
	}

	cfg := config.New(serveMux)

	normalMux.HandleFunc("POST /upload/{$}", handler.ExtractFile)

	go func() {
		if err := cfg.App.Server.ListenAndServe(); err != nil {
			cfg.App.Logger.Error(err.Error())
			os.Exit(1)
			return
		}
	}()

	conn, _ := net.Dial("udp", "1.1.1.1:80")
	localAddr := conn.LocalAddr().(*net.UDPAddr)
	conn.Close()

	cfg.App.Logger.Info("Server up and running", "IP", localAddr.IP.String(), "Port", "8080")

	select {}
}
