package api

import (
	"io/ioutil"
	"net/http"
	"sonnen-batterie-api/api/env"
)

func OpenApiDocumentation(env env.Environment, w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/yaml; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	dat, _ := ioutil.ReadFile("open-api-3.yaml")
	_, _ = w.Write(dat)
}
