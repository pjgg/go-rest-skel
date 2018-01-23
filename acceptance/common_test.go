// +build !unit

package acceptance

import (
	"bytes"
	"fmt"
	"io"
	"net/http"

	"github.com/DATA-DOG/godog"
)

type testContext struct {
	responseCode       int
	body               io.ReadCloser
	userName           string
	requestContentType string
	resp               *http.Response
}

func (context *testContext) resetContext(interface{}) {
	context.body = nil
	context.userName = ""
	context.responseCode = 0
	context.resp = nil
	context.requestContentType = ""
}

func (context *testContext) aContenttype(contentType string) error {
	context.requestContentType = contentType
	return nil
}

func (context *testContext) responseCodeMustBe(expectedResponseCode int) error {
	var e error
	if expectedResponseCode != context.responseCode {
		buf := new(bytes.Buffer)
		buf.ReadFrom(context.body)
		e = fmt.Errorf("expected response code to be: %d, but actual is: %d   .EXTRA INFO: Response body: %s", expectedResponseCode, context.responseCode, buf.String())
	}
	return e
}
func (context *testContext) responseMustBeAContentType(expectedContentType string) error {
	var e error
	var contentType = context.resp.Header.Get("Content-Type")

	if e != nil || contentType != expectedContentType {
		e = fmt.Errorf("expected content-type is: %s, but actual is: %s", expectedContentType, contentType)
	}
	return e
}

func CommonFeatureContext(commonSuit *godog.Suite, context *testContext) {
	commonSuit.Step(`^a content-type "([^"]*)"$`, context.aContenttype)
	commonSuit.Step(`^response code must be (\d+)$`, context.responseCodeMustBe)
	commonSuit.Step(`^response must be a content type "([^"]*)"$`, context.responseMustBeAContentType)
}
