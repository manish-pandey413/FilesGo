package main

import (
	"flag"
	"fmt"
	"log"
	"net"
	"net/http"

	"github.com/manish-pandey413/FilesGo/internal/handler"
)

func main() {

	webMode := flag.Bool("webMode", false, "Set true to use web mode.")
	path := flag.String("dir", "./", "Directory to share")
	flag.Parse()

	normalMux := http.NewServeMux()
	normalMux.HandleFunc("/upload/{$}", handler.ExtractFile)

	var (
		webModeMux *http.ServeMux
		fileServer http.Handler
	)
	if *webMode {
		webModeMux = http.NewServeMux()

		fileServer = http.FileServer(http.Dir(*path))
		webModeMux.Handle("GET /", fileServer)
	}

	handler := normalMux
	if *webMode {
		handler = webModeMux
	}

	server := &http.Server{
		Addr:    ":8080",
		Handler: handler,
	}

	go func() {
		if err := server.ListenAndServe(); err != nil {
			log.Fatalf("%s", err.Error())
			return
		}
	}()

	conn, _ := net.Dial("udp", "1.1.1.1:80")
	localAddr := conn.LocalAddr().(*net.UDPAddr)
	conn.Close()

	fmt.Printf("Running server on %s\n", localAddr.IP.String())

	select {}
}
