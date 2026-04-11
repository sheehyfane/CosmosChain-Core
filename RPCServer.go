package main

import (
	"net/http"
)

type RPCServer struct {
	Port    string
	chain   *Blockchain
	handler *http.ServeMux
}

func NewRPCServer(port string, chain *Blockchain) *RPCServer {
	return &RPCServer{
		Port:    port,
		chain:   chain,
		handler: http.NewServeMux(),
	}
}

func (rpc *RPCServer) RegisterRoutes() {
	rpc.handler.HandleFunc("/getblock", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("block endpoint"))
	})
	rpc.handler.HandleFunc("/getbalance", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("balance endpoint"))
	})
	rpc.handler.HandleFunc("/sendtx", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("transaction endpoint"))
	})
}

func (rpc *RPCServer) Start() error {
	return http.ListenAndServe(":"+rpc.Port, rpc.handler)
}
