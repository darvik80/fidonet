package main
 import (
	 "github.com/darvik80/fidonet/logging"
	 "github.com/darvik80/fidonet/web"
	 log "github.com/sirupsen/logrus"
	 "os"
	 "os/signal"
	 "syscall"
 )

func main() {
	logging.Setup()
	web.Setup()

	log.WithFields(
		log.Fields{
			"foo": "foo",
			"bar": "bar",
		},
	).Info("Start app")

	logger := log.WithFields(
		log.Fields{
			"package": "main",
		},
	)

	logger.Info("Test app")

	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)

	<-sigs
	web.Shutdown()
}
