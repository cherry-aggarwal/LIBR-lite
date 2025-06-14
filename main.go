package main

import (
	"log"
	"net/http"

	"github.com/cherry-aggarwal/libr/database"
	"github.com/cherry-aggarwal/libr/routers"
)

func main() {
	database.InitConnection()
	routers.Routers()
	log.Fatal(http.ListenAndServe(":3000", routers.Routers()))
	defer database.Pool.Close()

}
