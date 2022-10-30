package main

import (
	"net/http"

	handler "github.com/eltonCasacio/client-server-api/server/internal/handlers"
)

func main() {
	http.HandleFunc("/cotacao", handler.CotacaoHandler)
	http.ListenAndServe(":8080", nil)
}
