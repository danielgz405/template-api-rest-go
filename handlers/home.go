package handlers

import (
	"fmt"
	"net/http"

	"github.com/BacataCode/SmartCardConnectApiRest/pages"
	"github.com/BacataCode/SmartCardConnectApiRest/server"
	"github.com/gorilla/mux"
)

type HomeResponse struct {
	Message string `json:"message"`
	Status  int    `json:"status"`
}

func HomeHandler(s server.Server) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/html")

		params := mux.Vars(r)

		welcomePage, err := pages.Welcome(params["name"])
		if err != nil {
			fmt.Println(err.Error())
			return
		}

		w.WriteHeader(http.StatusOK)
		w.Write([]byte(welcomePage))
	}
}
