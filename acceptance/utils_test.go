// +build !unit

package acceptance

import (
	"bytes"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"os"
	"strings"
)

func makeHTTPQuery(req *http.Request) (*http.Response, error) {

	client := &http.Client{}
	resp, err := client.Do(req)

	return resp, err
}

func makeHTTPFormQuery(url string, data url.Values) (*http.Response, error) {

	client := &http.Client{}
	resp, err := client.PostForm(url, data)

	return resp, err
}

func newHTTPRequest(method, url string, body []byte, cookie *http.Cookie) (*http.Request, error) {
	req, err := http.NewRequest(method, url, bytes.NewBuffer(body))
	if nil != err {
		req = nil
	}

	if nil != cookie {
		req.AddCookie(cookie)
	}

	return req, err
}

func getCookie(w *http.Response, cookieName string) (string, error) {
	var err error
	header := w.Header.Get("Set-Cookie")

	parts := strings.Split(strings.TrimSpace(header), ";")

	cookie := strings.Split(parts[0], "=")

	if cookie[0] != cookieName {
		return "", errors.New("cookie not set")
	}

	return cookie[1], err
}

func getAddr() string {
	host, port := "localhost", "8080"

	if envHost, ok := os.LookupEnv("PORT_8080_TCP_ADDR"); ok {
		host = envHost
	}
	if envPort, ok := os.LookupEnv("PORT_8080_TCP_PORT"); ok {
		port = envPort
	}

	return fmt.Sprintf("%s:%s", host, port)
}
