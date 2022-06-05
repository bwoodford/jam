package server

import (
	"crypto/tls"
	"fmt"
	"io"
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

type Server struct {
	tls *tls.Config
}

func NewServer() (Server, error) {
	var server Server

	cer, err := tls.LoadX509KeyPair("server.crt", "server.key")
	if err != nil {
		log.Println(err)
		return server, err
	}

	config := &tls.Config{
		Certificates: []tls.Certificate{cer},
		MinVersion:   tls.VersionTLS13,
	}

	server = Server{
		tls: config,
	}

	return server, err
}

func (serv *Server) Start() {
	ln, err := tls.Listen("tcp", ":1965", serv.tls)
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

	msg, err := readRequestLine(conn)

	resp := handleRequest(string(msg), err)
	n, err := conn.Write(resp)

	if err != nil {
		log.Println(n, err)
		return
	}
}

func handleRequest(req string, err error) []byte {
	resp := SUCCESS
	if err != nil {
		resp += BAD_REQUEST
	} else {
		// deconstruct request
		println(req)
	}
	resp += "text/gemini; lang=en"
	return []byte(resp + "\r\n")
}

func readRequestLine(reader io.Reader) (line []byte, err error) {

	i := 0
	char := 0
	max := 1026
	line = make([]byte, 0, max)
	buff := make([]byte, 1)
	prevCarr := false

	for {
		char, err = reader.Read(buff)
		if i < max {
			if char > 0 {
				i++
				char := buff[0]
				if char == '\r' {
					prevCarr = true
					continue
				} else if char == '\n' && prevCarr {
					return
				} else if prevCarr {
					// new line was not found - add carriage return back
					prevCarr = false
					line = append(line, '\r')
				}
				line = append(line, char)
				continue
			}
		}
		return nil, fmt.Errorf("unable to parse client request")
	}
}
