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
	v1 "github.com/prometheus/client_golang/api/prometheus/v1"
	"github.com/prometheus/common/model"
	"github.com/prometheus/prometheus/model/timestamp"
	"github.com/prometheus/prometheus/prompb"
	"github.com/prometheus/prometheus/storage/remote"
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
	mockQuerier.EXPECT().Query(gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any()).DoAndReturn(
		func(_, _, _ any, _ ...any) (model.Vector, *v1.Warnings, error) {
			r := queryResult
			queryResult = model.Vector{} // return empty result on next call so it will exit the loop

			return r, nil, nil
		},
	).Times(2)

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

	err := client.ZeroMe(context.Background(), nowTime, metric)
	require.NoError(t, err)
}

func TestZeroMe_QueryError(t *testing.T) {
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
		queryErr = &v1.Error{Type: v1.ErrTimeout}
	)

	ctrl := gomock.NewController(t)
	t.Cleanup(ctrl.Finish)

	// setup querier mock
	mockQuerier := mock.NewMockAPI(ctrl)
	querier.SetV1API(mockQuerier)
	mockQuerier.EXPECT().Query(gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any()).Return(nil, nil, queryErr)

	// setup writer mock
	mockWriter := mock.NewMockWriteClient(ctrl)
	writer.SetClient(mockWriter)
	mockWriter.EXPECT().Store(gomock.Any(), gomock.Any(), 0).Return(nil).Times(0)

	client := New([]Metric{metric})

	err := client.ZeroMe(context.Background(), time.Now(), metric)
	require.ErrorIs(t, err, queryErr)
}

func TestZeroMe_EmptyResult(t *testing.T) {
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
	)

	ctrl := gomock.NewController(t)
	t.Cleanup(ctrl.Finish)

	// setup querier mock
	mockQuerier := mock.NewMockAPI(ctrl)
	querier.SetV1API(mockQuerier)
	mockQuerier.EXPECT().Query(gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any()).Return(model.Vector{}, nil, nil)

	// setup writer mock
	mockWriter := mock.NewMockWriteClient(ctrl)
	writer.SetClient(mockWriter)
	mockWriter.EXPECT().Store(gomock.Any(), gomock.Any(), 0).Return(nil).Times(0)

	client := New([]Metric{metric})

	err := client.ZeroMe(context.Background(), time.Now(), metric)
	require.NoError(t, err)
}

func TestZeroMe_WriteError(t *testing.T) {
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
		writeErr = &remote.HTTPError{}
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
		mockWriter.EXPECT().Store(gomock.Any(), gomock.Any(), 0).Return(writeErr)
	}

	client := New([]Metric{metric})

	err := client.ZeroMe(context.Background(), nowTime, metric)
	require.ErrorIs(t, err, writeErr)
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
