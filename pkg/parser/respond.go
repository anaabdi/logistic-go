package parser

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func Respond(w http.ResponseWriter, statusCode int, data interface{}) {
	//w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	b, err := json.Marshal(data)
	if err != nil {
		fmt.Printf("error when marshal: %s \n", err.Error())
	}

	w.Write(b)
}
