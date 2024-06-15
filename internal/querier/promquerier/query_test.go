package promquerier

import (
	"context"
	"testing"
	"time"

	"github.com/clwluvw/zerome/pkg/mock"
	v1 "github.com/prometheus/client_golang/api/prometheus/v1"
	"github.com/prometheus/common/model"
	"github.com/stretchr/testify/require"
	"go.uber.org/mock/gomock"
)

func newPromQuerier(t *testing.T) *PromQuerier {
	t.Helper()

	return &PromQuerier{}
}

func (pq *PromQuerier) mockV1API(t *testing.T) *mock.MockAPI {
	t.Helper()

	ctrl := gomock.NewController(t)
	t.Cleanup(ctrl.Finish)

	mockAPI := mock.NewMockAPI(ctrl)
	pq.v1api = mockAPI

	return mockAPI
}

func TestQuery(t *testing.T) {
	t.Parallel()

	var (
		metric         = "metric"
		interval       = time.Minute
		upLabels       = []string{"job", "instance"}
		expectedVector = model.Vector{
			&model.Sample{
				Metric: model.Metric{
					"job":      "test",
					"instance": "localhost:9090",
				},
				Value:     1,
				Timestamp: model.Now(),
			},
			&model.Sample{
				Metric: model.Metric{
					"job":      "test",
					"instance": "localhost:9091",
				},
				Value:     2,
				Timestamp: model.Now(),
			},
		}
	)

	queryTS := time.Now()

	promQuerier := newPromQuerier(t)
	promQuerier.mockV1API(t).EXPECT().Query(
		gomock.Any(),
		"count_over_time(metric[1m0s]) == 1 and metric and on(job,instance) (count_over_time(up[1m0s]) == 2)",
		queryTS,
		gomock.Any(),
	).Return(expectedVector, nil, nil)

	v, err := promQuerier.Query(context.Background(), queryTS, metric, interval, upLabels)
	require.NoError(t, err)
	require.Equal(t, expectedVector, v)
}

// TestQuery_WithWarnings tests the case when the Prometheus query returns warnings.
// The test ensures that the result is returned correctly even when there are warnings.
func TestQuery_WithWarnings(t *testing.T) {
	t.Parallel()

	var (
		metric         = "metric"
		interval       = time.Minute
		upLabels       = []string{"job", "instance"}
		expectedVector = model.Vector{
			&model.Sample{
				Metric: model.Metric{
					"job":      "test",
					"instance": "localhost:9090",
				},
				Value:     1,
				Timestamp: model.Now(),
			},
			&model.Sample{
				Metric: model.Metric{
					"job":      "test",
					"instance": "localhost:9091",
				},
				Value:     2,
				Timestamp: model.Now(),
			},
		}
	)

	queryTS := time.Now()

	promQuerier := newPromQuerier(t)
	promQuerier.mockV1API(t).EXPECT().Query(
		gomock.Any(),
		"count_over_time(metric[1m0s]) == 1 and metric and on(job,instance) (count_over_time(up[1m0s]) == 2)",
		queryTS,
		gomock.Any(),
	).Return(expectedVector, v1.Warnings{"warning!"}, nil)

	v, err := promQuerier.Query(context.Background(), queryTS, metric, interval, upLabels)
	require.NoError(t, err)
	require.Equal(t, expectedVector, v)
}

// TestQuery_UnexpectedResult tests the case when the Prometheus query returns an unexpected result.
func TestQuery_UnexpectedResult(t *testing.T) {
	t.Parallel()

	var (
		metric      = "metric"
		interval    = time.Minute
		upLabels    = []string{"job", "instance"}
		returnValue = model.Matrix{
			&model.SampleStream{
				Metric: model.Metric{
					"job":      "test",
					"instance": "localhost:9090",
				},
				Values: []model.SamplePair{
					{
						Timestamp: model.Now(),
						Value:     1,
					},
				},
			},
		}
	)

	queryTS := time.Now()

	promQuerier := newPromQuerier(t)
	promQuerier.mockV1API(t).EXPECT().Query(
		gomock.Any(),
		"count_over_time(metric[1m0s]) == 1 and metric and on(job,instance) (count_over_time(up[1m0s]) == 2)",
		queryTS,
		gomock.Any(),
	).Return(returnValue, nil, nil)

	v, err := promQuerier.Query(context.Background(), queryTS, metric, interval, upLabels)
	require.Error(t, err)
	require.Nil(t, v)
}

// TestQuery_WithErrors tests the case when the Prometheus query returns an error.
func TestQuery_WithErrors(t *testing.T) {
	t.Parallel()

	var (
		metric      = "metric"
		interval    = time.Minute
		upLabels    = []string{"job", "instance"}
		expectedErr = &v1.Error{Type: v1.ErrTimeout}
	)

	queryTS := time.Now()

	promQuerier := newPromQuerier(t)
	promQuerier.mockV1API(t).EXPECT().Query(
		gomock.Any(),
		"count_over_time(metric[1m0s]) == 1 and metric and on(job,instance) (count_over_time(up[1m0s]) == 2)",
		queryTS,
		gomock.Any(),
	).Return(nil, nil, expectedErr)

	v, err := promQuerier.Query(context.Background(), queryTS, metric, interval, upLabels)
	require.ErrorIs(t, err, expectedErr)
	require.Nil(t, v)
}
