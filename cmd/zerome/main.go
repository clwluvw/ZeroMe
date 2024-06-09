package main

import (
	"context"
	"log/slog"
	"os"
	"os/signal"
	"sync"
	"syscall"

	"github.com/clwluvw/zerome/internal/querier"
	"github.com/clwluvw/zerome/internal/querier/promquerier"
	"github.com/clwluvw/zerome/internal/writer"
	"github.com/clwluvw/zerome/internal/writer/promremotewriter"
	"github.com/clwluvw/zerome/internal/zerome"
)

//nolint:gochecknoglobals
var (
	Version   string
	Revision  string
	Branch    string
	BuildDate string
)

func buildQueriers(cfg map[string]zerome.QuerierConfig) map[string]querier.Querier {
	queriers := make(map[string]querier.Querier, len(cfg))

	for name, c := range cfg {
		switch c.Type {
		case "prometheus":
			q, err := promquerier.New(c.Address, c.Headers)
			if err != nil {
				panic(err)
			}

			queriers[name] = q
		default:
			panic("unknown querier type: " + c.Type)
		}
	}

	return queriers
}

func buildWriters(cfg map[string]zerome.WriterConfig) map[string]writer.Writer {
	writers := make(map[string]writer.Writer, len(cfg))

	for name, c := range cfg {
		switch c.Type {
		case "prometheus":
			w, err := promremotewriter.New(c.Address, c.Headers, c.Timeout)
			if err != nil {
				panic(err)
			}

			writers[name] = w
		default:
			panic("unknown writer type: " + c.Type)
		}
	}

	return writers
}

func main() {
	slog.Info("ZeroMe", "version", Version, "revision", Revision, "branch", Branch, "build_date", BuildDate)

	var cfg zerome.Config
	if err := zerome.LoadConfig("config.yaml", &cfg); err != nil {
		panic(err)
	}

	queriers := buildQueriers(cfg.Queriers)
	writers := buildWriters(cfg.Writers)

	metrics := make([]zerome.Metric, len(cfg.Metrics))
	for i, m := range cfg.Metrics {
		metrics[i] = m
		metrics[i].SetQuerier(queriers[m.Querier])
		metrics[i].SetWriter(writers[m.Writer])
	}

	zm := zerome.New(metrics)

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	var wg sync.WaitGroup

	wg.Add(1)

	zm.Run(ctx, &wg)

	// Wait for a signal to stop
	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, os.Interrupt, syscall.SIGTERM)

	go func() {
		<-signalChan
		cancel()
	}()

	wg.Wait()
}
