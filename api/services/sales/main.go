package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"runtime"
	"syscall"

	"github.com/mbpyytcqw/k8s/pkg/logger"
)

// build Defines for which enviroment curr. version of app.
var build string = "develop"

func main() {
	fmt.Println(1)
	ctx := context.Background()

	var log *logger.Logger

	events := logger.Events{
		Error: func(ctx context.Context, r logger.Record) {
			log.Info(ctx, "SEND ALERT")
		},
	}
	traceIDFn := func(ctx context.Context) string {
		return "" //web.GetTraceID()
	}

	log = logger.NewWithEvents(os.Stdout, logger.LevelInfo, "SALES", traceIDFn, events)

	_ = run(ctx, log)

}

func run(ctx context.Context, log *logger.Logger) error {
	log.Info(ctx, "startup server", "max procs", runtime.GOMAXPROCS(0), "enviroment", build)

	sigCh := make(chan os.Signal, 1)
	signal.Notify(sigCh, os.Interrupt, syscall.SIGTERM)
	sig := <-sigCh

	log.Info(ctx, "init shutdown", "signal", sig)
	defer log.Info(ctx, "shutdown complete", "signal", sig)

	return nil
}
