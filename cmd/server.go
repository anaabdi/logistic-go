package main

import (
	"context"
	"fmt"
	"net/http"

	"github.com/go-chi/chi"

	authctrl "logistic-go/api/auth/controller"
	pkgcfg "logistic-go/pkg/config"
	pkgconst "logistic-go/pkg/constant"
	pkgerr "logistic-go/pkg/error"
	"logistic-go/pkg/parser"
)

func main() {

	pkgcfg.InitConfig()

	router := chi.NewRouter()

	router.Route("/auth", func(r chi.Router) {
		// username + password
		r.With(ContentTypeMiddleware).Post("/login", authctrl.LoginHandler)
	})

	fmt.Printf("starting server %s on port %s \n", pkgcfg.AppConfig.AppName, pkgcfg.AppConfig.Port)

	targetHostServer := fmt.Sprintf("%s:%s", pkgcfg.AppConfig.Host, pkgcfg.AppConfig.Port)
	http.ListenAndServe(targetHostServer, router)
}

func ContentTypeMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		cType := r.Header.Get("Content-Type")

		fmt.Println("ctype: ", cType)

		if cType != "application/json" {
			parser.Respond(rw, http.StatusBadRequest, pkgerr.ErrorResponse{
				Code:    "BAD_REQ_CONTENT_TYPE",
				Message: "unexpected content type",
				Data:    nil,
			})
			return
		}

		ctx := r.Context()
		ctx = context.WithValue(ctx, pkgconst.CTYPE, cType)

		next.ServeHTTP(rw, r.WithContext(ctx))
	})
}
