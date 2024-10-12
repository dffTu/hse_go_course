package request

import (
	"net/http"
)

type Request struct {
	Handler func(http.ResponseWriter, *http.Request)
	Path    string
}
