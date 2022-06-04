package main

import (
	"bufio"
	"crypto/tls"
	"log"
	"net"
)

const (
	INPUT                             = "10 "
	INPUT_SENSITIVE                   = "11 "
	SUCCESS                           = "20 "
	REDIRECT_TEMPORARY                = "30 "
	REDIRECT_PERMANENT                = "31 "
	TEMPORARY_FAILURE                 = "40 "
	SERVER_UNAVAILABLE                = "41 "
	CGI_ERROR                         = "42 "
	PROXY_ERROR                       = "43 "
	SLOW_DOWN                         = "44 "
	PERMANENT_FAILURE                 = "50 "
	NOT_FOUND                         = "51 "
	GONE                              = "52 "
	PROXY_REQUEST_REFUSED             = "53 "
	BAD_REQUEST                       = "59 "
	CLIENT_CERTIFICATE_REQUIRED       = "60 "
	CLIENT_CERTIFICATE_NOT_AUTHORISED = "61 "
	CLIENT_CERTIFICATE_NOT_VALID      = "62 "
)

func main() {
	log.SetFlags(log.Lshortfile)

	cer, err := tls.LoadX509KeyPair("server.pem", "server.key")
	if err != nil {
		log.Println(err)
		return
	}

	config := &tls.Config{
		Certificates: []tls.Certificate{cer},
		MinVersion:   tls.VersionTLS13,
	}

	ln, err := tls.Listen("tcp", ":1965", config)
	if err != nil {
		log.Println(err)
		return
	}
	defer ln.Close()

	for {
		conn, err := ln.Accept()
		if err != nil {
			log.Println(err)
			continue
		}
		go handleConnection(conn)
	}
}

func handleConnection(conn net.Conn) {
	defer conn.Close()
	r := bufio.NewScanner(conn)

	resp := handleRequest(msg, err)
	n, err := conn.Write(resp)

	if err != nil {
		log.Println(n, err)
		return
	}
}

func handleRequest(req string, err error) []byte {
	resp := ""
	if err != nil {
		resp += BAD_REQUEST
	} else {
		// deconstruct request
	}
	return []byte(resp)
}
