package main

import (
	"testing"

	"github.com/clwluvw/zerome/internal/zerome"
	"github.com/stretchr/testify/require"
)

func TestBuildQueriers(t *testing.T) {
	t.Parallel()

	cfg := map[string]zerome.QuerierConfig{
		"prometheus": {
			Address: "http://localhost:9090",
			Headers: map[string]string{
				"Authorization": "Bearer token",
			},
			Type: "prometheus",
		},
		"prom2": {
			Address: "http://localhost:9091",
			Headers: map[string]string{
				"Authorization": "Bearer token",
			},
			Type: "prometheus",
		},
	}

	queriers := buildQueriers(cfg)
	for s := range cfg {
		require.Contains(t, queriers, s)
	}
}

func TestBuildQueriers_Unknown(t *testing.T) {
	t.Parallel()

	cfg := map[string]zerome.QuerierConfig{
		"prometheus": {
			Address: "http://localhost:9090",
			Headers: map[string]string{
				"Authorization": "Bearer token",
			},
			Type: "unknown",
		},
	}

	require.Panics(t, func() { buildQueriers(cfg) })
}

func TestBuildWriters(t *testing.T) {
	t.Parallel()

	cfg := map[string]zerome.WriterConfig{
		"prometheus": {
			Address: "http://localhost:9090",
			Headers: map[string]string{
				"Authorization": "Bearer token",
			},
			Timeout: 10,
			Type:    "prometheus",
		},
		"prom2": {
			Address: "http://localhost:9091",
			Headers: map[string]string{
				"Authorization": "Bearer token",
			},
			Timeout: 100,
			Type:    "prometheus",
		},
	}

	queriers := buildWriters(cfg)
	for s := range cfg {
		require.Contains(t, queriers, s)
	}
}

func TestBuildWriters_Unknown(t *testing.T) {
	t.Parallel()

	cfg := map[string]zerome.WriterConfig{
		"prometheus": {
			Address: "http://localhost:9090",
			Headers: map[string]string{
				"Authorization": "Bearer token",
			},
			Timeout: 10,
			Type:    "unknown",
		},
	}

	require.Panics(t, func() { buildWriters(cfg) })
}
