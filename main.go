package main

import (
	"context"
	"fmt"
	"log/slog"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/zarldev/gobase/config"
	"github.com/zarldev/gobase/ui"
)

func main() {
	slog.Info("starting...")
	startupInfo()
	ctx, cancel := context.WithCancel(context.Background())
	ui.StartUI(ctx, ui.Config{
		Port: config.ENVIRONMENT.PORT,
	})
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
	_, _ = fmt.Fprintf(os.Stdout, "\n")
	slog.Info("shutting down...")
	closeResources(cancel, sigs)
}

func closeResources(cancel context.CancelFunc, sigs chan os.Signal) {
	slog.Info("closing resources...")
	cancel()
	close(sigs)
	// wait for the server to shutdown
	time.Sleep(1 * time.Second)
}
