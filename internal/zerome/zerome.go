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

func (c *Client) Run(ctx context.Context, outerWG *sync.WaitGroup) {
	for _, metric := range c.metrics {
		go func(metric Metric, outerWG *sync.WaitGroup) {
			defer outerWG.Done()

			var wg sync.WaitGroup

			for {
				select {
				case <-ctx.Done():
					wg.Wait()

					return
				case <-time.After(metric.Interval):
					wg.Add(1)

					go func() {
						defer wg.Done()

						err := c.ZeroMe(ctx, time.Now(), metric)
						if err != nil {
							slog.ErrorContext(ctx, "Failed to zero metric", "metric", metric.Name, "error", err)
						}
					}()
				}
			}
		}(metric, outerWG)
	}
}

func (c *Client) ZeroMe(ctx context.Context, nowTime time.Time, metric Metric) error {
	// Query twice the interval to ensure that the metric has a missing data point in the past.
	queryInterval := metric.Interval * 2 //nolint:gomnd,mnd

	// Add query interval as a delay to cover exporter scrape failures.
	ts := nowTime.Add(-queryInterval)

	vector, err := metric.querier.Query(ctx, ts, metric.Name, queryInterval, metric.UpLabels)
	if err != nil {
		slog.ErrorContext(ctx, "Failed to query metric", "metric", metric.Name, "error", err)

		return err
	}

	timeSeries := c.zeroTimeSeries(metric, vector)
	if len(timeSeries) == 0 {
		return nil
	}

	err = metric.writer.Write(ctx, timeSeries, nil)
	if err != nil {
		slog.ErrorContext(ctx, "Failed to write metric", "metric", metric.Name, "error", err)

		return err
	}

	for _, s := range vector {
		slog.InfoContext(ctx, "Zeroed sample", "metric", metric.Name, "sample", s.String())
	}

	return c.ZeroMe(ctx, nowTime, metric) // Retry until nothing to zero
}

func (c *Client) zeroTimeSeries(metric Metric, vector model.Vector) []prompb.TimeSeries {
	timeSeries := make([]prompb.TimeSeries, 0, len(vector))

	for _, sample := range vector {
		timeSeries = append(timeSeries, prompb.TimeSeries{
			Labels: metricToLabels(metric.Name, sample.Metric),
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

func metricToLabels(metricName string, metric model.Metric) []prompb.Label {
	labels := make([]prompb.Label, 0, len(metric)+1)

	for k, v := range metric {
		labels = append(labels, prompb.Label{Name: string(k), Value: string(v)})
	}

	labels = append(labels, prompb.Label{Name: "__name__", Value: metricName})

	return labels
}
