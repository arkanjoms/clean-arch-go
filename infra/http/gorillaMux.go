package http

import (
	"clean-arch-go/infra/http/middleware"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
	"io"
	"net/http"
	"reflect"
	"strings"
)

type GorillaMux struct {
	Router *mux.Router
}

func NewGorillaMux() Http {
	router := mux.NewRouter().PathPrefix("/api").Subrouter()
	router.Use(middleware.NewCorsMiddleware("*").Apply)
	return GorillaMux{Router: router}
}

func (g GorillaMux) convertUrl(url string) string {
	return strings.ReplaceAll(url, "$", "")
}

func (g GorillaMux) On(method string, url string, handlerFn func(map[string]string, map[string][]string, io.ReadCloser) (interface{}, error)) {
	logrus.Debugf("Registering route => %s", url)
	g.Router.HandleFunc(g.convertUrl(url), func(writer http.ResponseWriter, request *http.Request) {
		pathParams := mux.Vars(request)
		queryParams := request.URL.Query()
		result, err := handlerFn(pathParams, queryParams, request.Body)
		if err != nil {
			http.Error(writer, err.Error(), http.StatusInternalServerError)
			return
		}
		if reflect.ValueOf(result).IsZero() {
			http.Error(writer, "not found", http.StatusNotFound)
			return
		}
		_ = json.NewEncoder(writer).Encode(result)
	}).Methods(strings.ToUpper(method))
}

func (g GorillaMux) Listen(port int) error {
	logrus.Infof("Server listening on port %d", port)
	return http.ListenAndServe(fmt.Sprint(":", port), g.Router)
}
