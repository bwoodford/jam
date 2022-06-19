package uri

import (
	"fmt"
	"log"
	"net/url"
)

type genericUri struct {
	scheme        string
	authorization string
	path          string
	fragment      string
}

type Uri struct {
	Scheme   string
	Host     string
	Port     string
	Path     string
	Fragment string
	Full     string
}

func Normalize(req string) (Uri, error) {
	uri, err := newGeminiUri(req)
	if err != nil {
		return Uri{}, err
	}
	if !uri.validate() {
		return Uri{}, fmt.Errorf("invalid gemini uri")
	}
	return uri, nil
}

func newGeminiUri(req string) (Uri, error) {
	var gem Uri

	u, err := url.Parse(req)
	if err != nil {
		log.Fatal(err)
		return Uri{}, err
	}

	gem.Fragment = u.Fragment
	gem.Scheme = u.Scheme
	gem.Host = u.Hostname()
	gem.Port = u.Port()
	gem.Path = u.Path
	gem.Full = u.String()

	if gem.Port == "" {
		gem.Port = "1965"
	}

	if gem.Path == "" {
		gem.Path = "/"
	}

	return gem, nil
}

func (uri *Uri) validate() (valid bool) {
	valid = true
	if uri.Host == "" ||
		uri.Scheme == "" {
		valid = false
	}
	return
}
