package main

import (
	"net/http"

	"github.com/AJackTi/go-template-fx/httphandler"
	"go.uber.org/zap"
)

func main() {
	logger, _ := zap.NewProduction()
	defer logger.Sync()
	slogger := logger.Sugar()

	mux := http.NewServeMux()
	httphandler.New(mux, slogger)

	http.ListenAndServe(":8080", mux)
}
