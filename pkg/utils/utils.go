package utils

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)


func ParseBody(r *http.Request, x interface{}) interface{} {
	body, errR := ioutil.ReadAll(r.Body);
	if errR != nil {
		log.Fatal(errR)
	}
	errM := json.Unmarshal(body, x)
	if errM != nil {
		log.Fatal(errM)
	}
	return x
}