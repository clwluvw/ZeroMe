package zerome

import (
	"context"
	"testing"
	"time"

	"github.com/clwluvw/zerome/internal/querier/promquerier"
	"github.com/clwluvw/zerome/internal/writer/promremotewriter"
	"github.com/clwluvw/zerome/pkg/mock"
	"github.com/gogo/protobuf/proto"
	"github.com/golang/snappy"
	"github.com/prometheus/common/model"
	"github.com/prometheus/prometheus/model/timestamp"
	"github.com/prometheus/prometheus/prompb"
	"github.com/stretchr/testify/require"
	"go.uber.org/mock/gomock"
)

func TestZeroMe(t *testing.T) {
	t.Parallel()

	var (
		querier = &promquerier.PromQuerier{}
		writer  = &promremotewriter.PromRemoteWrite{}
		metric  = Metric{
			Name:     "metric",
			Interval: time.Minute,
			UpLabels: []string{"job", "instance"},
			querier:  querier,
			writer:   writer,
		}
		nowTime     = time.Now()
		queryResult = model.Vector{
			&model.Sample{
				Metric: model.Metric{
					"job": "test",
				},
				Value:     1,
				Timestamp: model.Time(timestamp.FromTime(nowTime)),
			},
		}
		expectedTimeSeries = []prompb.TimeSeries{
			{
				Labels: []prompb.Label{
					{
						Name:  "job",
						Value: "test",
					},
					{
						Name:  "__name__",
						Value: "metric",
					},
				},
				Samples: []prompb.Sample{
					{
						Timestamp: timestamp.FromTime(nowTime.Add(-metric.Interval)),
						Value:     0,
					},
				},
			},
		}
		expectedWriteRequest = &prompb.WriteRequest{
			Timeseries: expectedTimeSeries,
			Metadata:   nil,
		}
	)

	ctrl := gomock.NewController(t)
	t.Cleanup(ctrl.Finish)

	// setup querier mock
	mockQuerier := mock.NewMockAPI(ctrl)
	querier.SetV1API(mockQuerier)
	mockQuerier.EXPECT().Query(gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any()).Return(queryResult, nil, nil)

	// setup writer mock
	{
		mockWriter := mock.NewMockWriteClient(ctrl)
		writer.SetClient(mockWriter)

		pBuf := proto.NewBuffer(nil)
		err := pBuf.Marshal(expectedWriteRequest)
		require.NoError(t, err)

		snappyBytes := snappy.Encode(nil, pBuf.Bytes())
		mockWriter.EXPECT().Store(gomock.Any(), snappyBytes, 0).Return(nil)
	}

	client := New([]Metric{metric})

	err := client.ZeroMe(context.Background(), metric)
	require.NoError(t, err)
}

func TestZeroTimeSeries(t *testing.T) {
	t.Parallel()

	var (
		metric = Metric{
			Name:     "metric",
			Interval: time.Minute,
			UpLabels: []string{"job", "instance"},
		}
		nowTime     = time.Now()
		queryResult = model.Vector{
			&model.Sample{
				Metric: model.Metric{
					"job": "test",
				},
				Value:     1,
				Timestamp: model.Time(timestamp.FromTime(nowTime)),
			},
		}
		expectedTimeSeries = []prompb.TimeSeries{
			{
				Labels: []prompb.Label{
					{
						Name:  "job",
						Value: "test",
					},
					{
						Name:  "__name__",
						Value: "metric",
					},
				},
				Samples: []prompb.Sample{
					{
						Timestamp: timestamp.FromTime(nowTime.Add(-metric.Interval)),
						Value:     0,
					},
				},
			},
		}
	)

	client := New([]Metric{metric})

	timeSeries := client.zeroTimeSeries(metric, queryResult)
	require.Equal(t, expectedTimeSeries, timeSeries)
}

func TestLabels(t *testing.T) {
	t.Parallel()

	var (
		metricName    = "metric"
		samplesMetric = model.Metric{
			"job": "test",
		}
		expectedLabels = []prompb.Label{
			{
				Name:  "job",
				Value: "test",
			},
			{
				Name:  "__name__",
				Value: "metric",
			},
		}
	)

	labels := metricToLabels(metricName, samplesMetric)
	require.Equal(t, expectedLabels, labels)
}
