package logging

import log "github.com/sirupsen/logrus"


func Setup() {
	log.SetFormatter(&log.TextFormatter{
		FullTimestamp: true,
	})
}