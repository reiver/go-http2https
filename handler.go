package http2https

import (
	"net/http"
)

// Handler is the default [http.Handler] for this package.
// Handler does not do any logging.
//
// You can use it rather than creating a new [HTTPHandler] if you do not need any logging.
//
// If you need logging, then create your own [HTTPHandler] instead.
var Handler http.Handler = HTTPHandler{}
