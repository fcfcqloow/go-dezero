package util

import (
	"os"
	"os/signal"
)

func WaitCtrlC() {
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	<-quit

}
