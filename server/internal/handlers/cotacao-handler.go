package handler

import (
	"context"
	"encoding/json"
	"net/http"
	"time"

	"github.com/eltonCasacio/client-server-api/server/internal/database"
	"github.com/eltonCasacio/client-server-api/server/pkg"
)

func CotacaoHandler(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := context.WithTimeout(r.Context(), time.Millisecond*300)
	defer cancel()

	select {
	case <-ctx.Done():
		w.WriteHeader(http.StatusRequestTimeout)
	default:
		cotacao, err := pkg.BuscaCotacao()
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		err = database.SalvarCotacaoSqlite(*cotacao)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(cotacao.Bid)
	}
}
