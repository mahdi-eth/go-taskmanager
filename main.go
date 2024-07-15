package main

import (
	"log"
	"net/http"
	"runtime"

	"github.com/codegangsta/negroni"
	common "github.com/mahdi-eth/go-taskmanager/common"
	"github.com/mahdi-eth/go-taskmanager/routers"
)

func main() {
	runtime.GOMAXPROCS(1)

	router := routers.InitRoutes()
	n := negroni.Classic()
	n.UseHandler(router)

	server := &http.Server{
		Addr:    common.AppConfig.Server,
		Handler: n,
	}
	log.Println("Listening...")
	server.ListenAndServe()
}
