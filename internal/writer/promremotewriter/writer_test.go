package promremotewriter

import (
	"context"
	"testing"

	"github.com/clwluvw/zerome/pkg/mock"
	"github.com/gogo/protobuf/proto"
	"github.com/golang/snappy"
	"github.com/prometheus/prometheus/prompb"
	"github.com/stretchr/testify/require"
	"go.uber.org/mock/gomock"
)

func newPromRemoteWrite(t *testing.T) *PromRemoteWrite {
	t.Helper()

	return &PromRemoteWrite{}
}

func (prw *PromRemoteWrite) mockClient(t *testing.T) *mock.MockWriteClient {
	t.Helper()

	ctrl := gomock.NewController(t)
	t.Cleanup(ctrl.Finish)

	mockWriteClient := mock.NewMockWriteClient(ctrl)
	prw.client = mockWriteClient

	return mockWriteClient
}

func TestWrite(t *testing.T) {
	t.Parallel()

	var (
		timeSeries = []prompb.TimeSeries{
			{
				Labels: []prompb.Label{
					{
						Name:  "job",
						Value: "test",
					},
					{
						Name:  "instance",
						Value: "localhost:9090",
					},
				},
				Samples: []prompb.Sample{
					{
						Timestamp: 1,
						Value:     1,
					},
				},
			},
		}
		metadata = []prompb.MetricMetadata{
			{
				Type:             prompb.MetricMetadata_COUNTER,
				MetricFamilyName: "test",
				Help:             "test",
			},
		}
	)

	writeRequest := &prompb.WriteRequest{
		Timeseries: timeSeries,
		Metadata:   metadata,
	}

	pBuf := proto.NewBuffer(nil)
	err := pBuf.Marshal(writeRequest)
	require.NoError(t, err)

	snappyBytes := snappy.Encode(nil, pBuf.Bytes())

	promRemoteWrite := newPromRemoteWrite(t)
	promRemoteWrite.mockClient(t).EXPECT().Store(
		gomock.Any(),
		snappyBytes,
		0,
	).Return(nil)

	err = promRemoteWrite.Write(context.Background(), timeSeries, metadata)
	require.NoError(t, err)
}
