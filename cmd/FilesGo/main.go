package main

import (
	"fmt"
	"log"
	"net"
	"net/http"

	"github.com/manish-pandey413/FilesGo/internal/handler"
)

func main() {
	mux := http.NewServeMux()

	mux.Handle("/upload/{$}", http.HandlerFunc(handler.ExtractFile))

	go func() {
		if err := http.ListenAndServe(":8080", mux); err != nil {
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
