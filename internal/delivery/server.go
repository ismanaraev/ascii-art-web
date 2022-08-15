package delivery

import (
	"ascii-art-web/internal/delivery/handlers"
	"ascii-art-web/internal/entity"
	"fmt"
	"net/http"
	"os"
)

//In this function we take interfaces for generator as an argument, then init handlers, passing interfaces as an argument for generating ascii-art, then register the handlers and start server
func StartServer(generator entity.AsciiArtGenerator) error {
	router := http.NewServeMux()
	fs := http.FileServer(http.Dir("static"))
	handler := handlers.NewHandler(generator)
	router.Handle("/static/", http.StripPrefix("/static/", fs))
	router.HandleFunc("/", handler.AsciiArtMainPage)
	router.HandleFunc("/ascii-art", handler.AsciiArtReadyPage)
	router.HandleFunc("/api", handler.AsciiAPI)
	router.HandleFunc("/download", handler.AsciiArtDownloadFile)
	port := os.Getenv("ASCII_WEB_PORT")
	fmt.Printf("stating server at %s\n", port)
	server := http.Server{
		Handler: router,
		Addr:    ":" + port,
	}
	return server.ListenAndServe()
}
