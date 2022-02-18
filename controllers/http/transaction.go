package http

import (
	"context"
	"net/http"

	"github.com/etheralley/etheralley-core-api/common"
	"github.com/etheralley/etheralley-core-api/usecases"
)

func (hc *HttpController) parseTransaction(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		query := r.URL.Query()
		ctx := r.Context()

		tx := &usecases.TransactionInput{
			Id:         query.Get("tx_id"),
			Blockchain: query.Get("blockchain"),
		}

		ctx = context.WithValue(ctx, common.ContextKeyTransaction, tx)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func (hc *HttpController) getInteractionRoute(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	query := r.URL.Query()

	interaction, err := hc.getInteraction(ctx, &usecases.GetInteractionInput{
		Address: query.Get("user_address"),
		Interaction: &usecases.InteractionInput{
			Transaction: ctx.Value(common.ContextKeyTransaction).(*usecases.TransactionInput),
			Type:        query.Get("type"),
		},
	})

	if err != nil {
		hc.presenter.PresentBadRequest(ctx, w, err)
		return
	}

	hc.presenter.PresentInteraction(ctx, w, interaction)
}
