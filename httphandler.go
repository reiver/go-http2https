package http2https

import (
	"net/http"
	"net/url"

	"github.com/reiver/go-http500"
)

// HTTPHandler deals with redirecting a HTTP request to an HTTPS request.
//
// So, for example, a request to:
//
//	http://example.com/apple/banana/cherry
//
// Would get redirected to:
//
//	https://example.com/apple/banana/cherry
//
// (Note that the scheme of the URL changed from "http" to "https".)
type HTTPHandler struct {
	LogError func(...any)
}

var _ http.Handler = HTTPHandler{}

func (receiver HTTPHandler) ServeHTTP(responsewriter http.ResponseWriter, request *http.Request) {
	var logError func(...any) = receiver.LogError
	if nil == logError {
		logError = discard
	}

	if nil == responsewriter {
		go logError("nil http-response-writer")
		return
	}

	if nil == request {
		http500.InternalServerError(responsewriter, request)
		go logError("nil http-request")
		return
	}
	if nil == request.URL {
		http500.InternalServerError(responsewriter, request)
		go logError("nil http-request-url")
		return
	}

	var redirecturi string
	{
		var redirectto url.URL = *request.URL
		redirectto.Scheme = "https"
		if "" == redirectto.Host {
			redirectto.Host = request.Host
		}

		redirecturi = redirectto.String()
	}

	http.Redirect(responsewriter, request, redirecturi, http.StatusTemporaryRedirect)
}
