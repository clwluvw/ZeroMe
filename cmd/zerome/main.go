package main

import (
	"context"
	"fmt"
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

func buildQueriers(cfg map[string]zerome.QuerierConfig) (map[string]querier.Querier, error) {
	queriers := make(map[string]querier.Querier, len(cfg))

	for name, c := range cfg {
		switch c.Type {
		case "prometheus":
			q, err := promquerier.New(c.Address, c.Headers)
			if err != nil {
				return nil, err
			}

			queriers[name] = q
		default:
			return nil, fmt.Errorf("unknown querier type: %s", c.Type)
		}
	}

	return queriers, nil
}

func buildWriters(cfg map[string]zerome.WriterConfig) (map[string]writer.Writer, error) {
	writers := make(map[string]writer.Writer, len(cfg))

	for name, c := range cfg {
		switch c.Type {
		case "prometheus":
			w, err := promremotewriter.New(c.Address, c.Headers, c.Timeout)
			if err != nil {
				return nil, err
			}

			writers[name] = w
		default:
			return nil, fmt.Errorf("unknown writer type: %s", c.Type)
		}
	}

	return writers, nil
}

func main() {
	var cfg zerome.Config
	if err := zerome.LoadConfig("config.yaml", &cfg); err != nil {
		panic(err)
	}

	queriers, err := buildQueriers(cfg.Queriers)
	if err != nil {
		panic(err)
	}

	writers, err := buildWriters(cfg.Writers)
	if err != nil {
		panic(err)
	}

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
