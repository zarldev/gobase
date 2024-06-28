package main

import (
	"context"
	"fmt"
	"log/slog"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/zarldev/go-base/config"
	"github.com/zarldev/go-base/ui"
)

func main() {
	slog.Info("starting...")
	startupInfo()
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	ui.StartUI(ctx, ui.Config{Port: config.ENVIRONMENT.PORT})
	slog.Info("started.")
	waitForInterrupt(cancel)
	slog.Info("shutdown.")
}

func startupInfo() {
	slog.Info("build information:",
		slog.String("name", config.ENVIRONMENT.NAME),
		slog.String("version", config.ENVIRONMENT.VERSION),
		slog.String("environment", config.ENVIRONMENT.ENV),
		slog.String("build_tags", config.ENVIRONMENT.BUILD_TAGS),
	)
	slog.Info("runtime information:",
		slog.String("log_level", config.ENVIRONMENT.LOG_LEVEL.String()),
		slog.String("port", config.ENVIRONMENT.PORT),
	)
}

func waitForInterrupt(cancel context.CancelFunc) {
	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT)
	slog.Info("press ctrl+c to shutdown...")
	<-sigs
	slog.Info("shutting down...")
	closeResources(cancel, sigs)
}

func closeResources(cancel context.CancelFunc, sigs chan os.Signal) {
	slog.Info("closing resources...")
	_, _ = fmt.Fprintln(os.Stdout, "")
	cancel()
	close(sigs)
	// wait for the server to shutdown
	time.Sleep(1 * time.Second)
}
