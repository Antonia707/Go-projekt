package main

import (
	"fmt"
	"net/http"

	"Projekt/config"
	"Projekt/handlers"
	"Projekt/middleware"
)

func main() {
	http.HandleFunc("/jmbag", handlers.JmbagHandler)
	http.HandleFunc("/sum", middleware.Authenticate(handlers.SumHandler))
	http.HandleFunc("/fetch", handlers.FetchHandler)
	http.HandleFunc("/0036540336", middleware.Authenticate(handlers.MyJmbagHandler))
	http.HandleFunc("/saxpy", middleware.Authenticate(handlers.SaxpyHandler))

	conf, err := config.GetConfigFromYaml()
	if err == nil {
		err1 := http.ListenAndServe(conf.HTTPConf.Address+":"+conf.HTTPConf.Port, nil)
		if err1 != nil {
			fmt.Println("Server error:", err)
		}
	} else {
		fmt.Println("Server error:", err)
	}
}
