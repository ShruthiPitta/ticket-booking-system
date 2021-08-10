package app

import (
	"fmt"
	"net/http"
)

func StartServer() {
	server := &http.Server{
		Addr: ":8080",
	}

	http.HandleFunc("/book", bookTicketHandler)
	http.HandleFunc("/", pingHandler())

	if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		fmt.Errorf("something went wrong %s", err)
		fmt.Println("Server Not Started")
		return
	}
	fmt.Println("Server Started")
}

func bookTicketHandler(writer http.ResponseWriter, request *http.Request) {
	writer.WriteHeader(http.StatusCreated)
	fmt.Fprintf(writer, "Ticket Created!")
}

func pingHandler() http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		writer.WriteHeader(http.StatusOK)
	}
}
