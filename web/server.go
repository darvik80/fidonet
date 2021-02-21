package web

import (
	"context"
	"github.com/darvik80/fidonet/web/rpc"
	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
	"net/http"
	"time"
)

var srv *http.Server
var log = logrus.WithFields(
	logrus.Fields{
		"package":   "web",
		"component": "server",
	},
)

func Setup() {
	log.Info("Setup WEB Server")
	r := mux.NewRouter()

	r.HandleFunc("/", func(writer http.ResponseWriter, req *http.Request) {
		writer.Write([]byte("<H1>Hello World</H1>"))
	})

	if h, err := rpc.Setup(); err == nil {
		r.Handle("/rpc", h)
	}

	srv = &http.Server{
		Handler: r,
		Addr:    "127.0.0.1:8000",
		// Good practice: enforce timeouts for servers you create!
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	go func() {
		if err := srv.ListenAndServe(); err != nil {
			log.Errorf("Failed setup WEB server, %s", err)
		}
	}()
}

func Shutdown() {
	log.Info("Shutdown WEB Server")
	// Create a deadline to wait for.
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*15)
	defer cancel()
	// Doesn't block if no connections, but will otherwise wait
	// until the timeout deadline.
	if err := srv.Shutdown(ctx); err != nil {
		log.Errorf("Failed shutdow WEB server, %s", err)
	}

	if err := srv.Close(); err != nil {
		log.Errorf("Failed shutdow WEB server, %s", err)
	}
}
