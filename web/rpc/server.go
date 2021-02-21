package rpc

import (
	"github.com/gorilla/rpc/v2"
	"github.com/gorilla/rpc/v2/json"
	log "github.com/sirupsen/logrus"
)

func Setup() (*rpc.Server, error) {
	s := rpc.NewServer()
	s.RegisterCodec(json.NewCodec(), "application/json")
	s.RegisterCodec(json.NewCodec(), "application/json;charset=UTF-8")
	if err := s.RegisterService(new(Health), ""); err != nil {
		log.Errorf("Failed register RPC service, %s", err)
		return nil, err
	}

	return s, nil
}
