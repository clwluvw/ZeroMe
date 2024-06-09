package writer

import (
	"context"

	"github.com/prometheus/prometheus/prompb"
)

type Writer interface {
	Write(ctx context.Context, timeSeries []prompb.TimeSeries, metadata []prompb.MetricMetadata) error
}
