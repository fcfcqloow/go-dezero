package sysutil

import (
	"net/http"
	_ "net/http/pprof"
	"runtime"

	"github.com/fcfcqloow/go-advance/log"
)

func Stats() (memStats runtime.MemStats) {
	runtime.ReadMemStats(&memStats)
	return
}

func RunViewMemoryLocal() {
	go func() {
		log.Info("http://localhost:6060/debug/pprof")
		log.Info(http.ListenAndServe("localhost:6060", nil))
	}()
}
