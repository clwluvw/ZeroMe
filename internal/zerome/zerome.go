package zerome

import (
	"context"
	"log/slog"
	"sync"
	"time"

	"github.com/prometheus/common/model"
	"github.com/prometheus/prometheus/model/timestamp"
	"github.com/prometheus/prometheus/prompb"
)

type Client struct {
	metrics []Metric
}

func New(metrics []Metric) *Client {
	return &Client{
		metrics: metrics,
	}
}

func (c *Client) Run(ctx context.Context, wg *sync.WaitGroup) {
	for _, metric := range c.metrics {
		go func(metric Metric, wg *sync.WaitGroup) {
			defer wg.Done()

			for {
				select {
				case <-ctx.Done():
					return
				case <-time.After(metric.Interval):
					err := c.ZeroMe(ctx, metric)
					if err != nil {
						slog.ErrorContext(ctx, "Failed to zero metric", "metric", metric.Name, "error", err)
					}
				}
			}
		}(metric, wg)
	}
}

func (c *Client) ZeroMe(ctx context.Context, metric Metric) error {
	vector, err := metric.querier.Query(ctx, metric.Name, metric.Interval*2, metric.UpLabels)
	if err != nil {
		slog.ErrorContext(ctx, "Failed to query metric", "metric", metric.Name, "error", err)
		return err
	}

	timeSeries := c.zeroMetric(metric, vector)
	if len(timeSeries) == 0 {
		return nil
	}

	err = metric.writer.Write(ctx, timeSeries, nil)
	if err != nil {
		slog.ErrorContext(ctx, "Failed to write metric", "metric", metric.Name, "error", err)
		return err
	}

	slog.InfoContext(ctx, "Zeroed metric", "metric", metric.Name, "vector", vector)

	return nil
}

func (c *Client) zeroMetric(metric Metric, vector model.Vector) []prompb.TimeSeries {
	timeSeries := make([]prompb.TimeSeries, 0, len(vector))

	for _, sample := range vector {
		timeSeries = append(timeSeries, prompb.TimeSeries{
			Labels: c.labels(metric.Name, sample.Metric),
			Samples: []prompb.Sample{
				{
					Timestamp: timestamp.FromTime(sample.Timestamp.Time().Add(-metric.Interval)),
					Value:     0,
				},
			},
		})
	}

	return timeSeries
}

func (c *Client) labels(metricName string, metric model.Metric) []prompb.Label {
	labels := make([]prompb.Label, 0, len(metric)+1)

	for k, v := range metric {
		labels = append(labels, prompb.Label{Name: string(k), Value: string(v)})
	}

	labels = append(labels, prompb.Label{Name: "__name__", Value: metricName})

	return labels
}
