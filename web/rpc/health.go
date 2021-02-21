package rpc

import (
	log "github.com/sirupsen/logrus"
	"net/http"
)

type Args struct {

}

type Result struct {
	Status string `json:"status"`
}

type Health struct{}

func (h *Health) Handle(r *http.Request, args* Args, result *Result) error {
	log.Info("Handle Health")
	*result = Result{
		Status: "UP",
	}

	return nil
}