package http

import "io"

type Http interface {
	On(string, string, func(map[string]string, map[string][]string, io.ReadCloser) (interface{}, error))
	Listen(int) error
}
