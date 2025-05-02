package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/HakashiKatake/Students-Rest-Api-Go/internal/config"
)

func main() {
	//load config
	cfg := config.MustLoad()
	//db setup

	//setup router
	router := http.NewServeMux()
	router.HandleFunc("GET /", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello, World!"))
	})

	//setup server
	server := http.Server{
		Addr:    cfg.Addr,
		Handler: router,
	}
	fmt.Println("server started")

	err := server.ListenAndServe()
	if err != nil {
		log.Fatal("failed to start server: ", err)
	}

	fmt.Println("server started on ", cfg.Addr)

}
