package controller

import (
	"encoding/json"
	"fmt"
	pkgcfg "logistic-go/pkg/config"
	pkgconst "logistic-go/pkg/constant"
	"logistic-go/pkg/parser"
	"net/http"
)

type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

// {
// 	"code": "BAD_REQUEST_PARSING"
// 	"message": "failed to parse"
// 	"data":
// }

type LoginResponse struct {
	Code    string      `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func LoginHandler(w http.ResponseWriter, r *http.Request) {

	var loginRequest LoginRequest
	fmt.Println(loginRequest)

	fmt.Println("env appname: ", pkgcfg.AppConfig.AppName)

	ctype := r.Context().Value(pkgconst.CTYPE)
	fmt.Println("ctype: ", ctype)

	err := json.NewDecoder(r.Body).Decode(&loginRequest)
	if err != nil {
		fmt.Printf("error here %s \n", err.Error())

		response := LoginResponse{
			Code:    "BAD_REQUEST_PARSING",
			Message: "failed to parse the req body",
			Data:    nil,
		}

		parser.Respond(w, http.StatusBadRequest, response)
		return
	}

	parser.Respond(w, http.StatusOK, nil)
	fmt.Println(loginRequest)
}
