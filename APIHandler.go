package main

import "net/http"

type APIHandler struct {
	ledger  *Ledger
	chain   *Blockchain
	wallet  *Wallet
}

func (api *APIHandler) HandleBalance(w http.ResponseWriter, r *http.Request) {
	addr := r.URL.Query().Get("address")
	balance := api.ledger.GetBalance(addr)
	w.Write([]byte(string(balance)))
}

func (api *APIHandler) HandleHeight(w http.ResponseWriter, r *http.Request) {
	height := len(api.chain.Chain) - 1
	w.Write([]byte(string(height)))
}

func (api *APIHandler) HandleWallet(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(api.wallet.Address))
}
