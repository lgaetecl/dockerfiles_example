package handler

import (
	"bytes"
	"io/ioutil"
	"log"
	"net/http"
)

// BodyLogging ...
func BodyLogging(w http.ResponseWriter, r *http.Request) {

	body, _ := ioutil.ReadAll(r.Body)
	log.Printf("%s\t%s\t%s\t%v",
		r.RemoteAddr,
		r.Method,
		r.RequestURI,
		ioutil.NopCloser(bytes.NewBuffer(body)),
	)
}
