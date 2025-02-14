package signals

import (
	"os"
	"os/signal"
	"syscall"

	log "github.com/sirupsen/logrus"
)

func CatchSigterm(isDone chan struct{}) {
	signals := make(chan os.Signal, 1)
	signal.Notify(signals, syscall.SIGTERM, syscall.SIGINT)

	go func() {
		sig := <-signals
		log.Debugf("Got system signal: %s", sig.String())
		isDone <- struct{}{}
	}()
}
