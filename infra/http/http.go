package http

import "io"

type Http interface {
	On(method string, url string, handlerFn func(map[string]string, map[string][]string, io.ReadCloser) (interface{}, error))
	Listen(port int) error
}
