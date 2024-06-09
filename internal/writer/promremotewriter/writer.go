package promremotewriter

import (
	"context"
	"net/url"
	"time"

	"github.com/clwluvw/zerome/internal/writer"
	"github.com/gogo/protobuf/proto"
	"github.com/golang/snappy"
	"github.com/prometheus/common/config"
	"github.com/prometheus/common/model"
	"github.com/prometheus/prometheus/prompb"
	"github.com/prometheus/prometheus/storage/remote"
)

var _ writer.Writer = &PromRemoteWrite{}

type PromRemoteWrite struct {
	client remote.WriteClient
}

func (prw *PromRemoteWrite) SetClient(client remote.WriteClient) {
	prw.client = client
}

func New(address string, headers map[string]string, timeout time.Duration) (*PromRemoteWrite, error) {
	u, err := url.Parse(address)
	if err != nil {
		return nil, err
	}

	if headers == nil {
		headers = make(map[string]string, 4)
	}

	headers["Content-Type"] = "application/x-protobuf"
	headers["Content-Encoding"] = "snappy"
	headers["User-Agent"] = "ZeroMe"
	headers["X-Prometheus-Remote-Write-Version"] = "0.1.0"

	client, err := remote.NewWriteClient("ZeroMe", &remote.ClientConfig{
		URL:     &config.URL{URL: u},
		Headers: headers,
		Timeout: model.Duration(timeout),
	})
	if err != nil {
		return nil, err
	}

	return &PromRemoteWrite{
		client: client,
	}, nil
}

func (prw *PromRemoteWrite) Write(ctx context.Context, timeSeries []prompb.TimeSeries, metadata []prompb.MetricMetadata) error {
	writeReq := &prompb.WriteRequest{
		Timeseries: timeSeries,
		Metadata:   metadata,
	}

	pBuf := proto.NewBuffer(nil)
	err := pBuf.Marshal(writeReq)
	if err != nil {
		return err
	}

	snappyBytes := snappy.Encode(nil, pBuf.Bytes())

	return prw.client.Store(ctx, snappyBytes, 0)
}
