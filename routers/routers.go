package routers

import (
	"fmt"

	"github.com/cherry-aggarwal/libr/controller"
	"github.com/gorilla/mux"
)

func Routers() *mux.Router {
	fmt.Println("setting up the routers")

	R := mux.NewRouter()
	R.HandleFunc("/", controller.HomeHandler).Methods("GET")
	R.HandleFunc("/submit", controller.MsgIN).Methods("POST")
	R.HandleFunc("/fetch/{timestamp}", controller.MsgOUT).Methods("GET")

	return R
}
