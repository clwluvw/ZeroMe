package promquerier

import (
	"context"
	"fmt"
	"log/slog"
	"strings"
	"time"

	"github.com/clwluvw/zerome/internal/querier"
	"github.com/prometheus/client_golang/api"
	v1 "github.com/prometheus/client_golang/api/prometheus/v1"
	"github.com/prometheus/common/model"
)

var _ querier.Querier = &PromQuerier{}

type PromQuerier struct {
	v1api v1.API
}

func (pq *PromQuerier) SetV1API(api v1.API) {
	pq.v1api = api
}

func New(address string, headers map[string]string) (*PromQuerier, error) {
	client, err := api.NewClient(api.Config{
		Address:      address,
		RoundTripper: &headerTransport{headers: headers},
	})
	if err != nil {
		return nil, err
	}

	return &PromQuerier{
		v1api: v1.NewAPI(client),
	}, nil
}

func (pq *PromQuerier) Query(ctx context.Context, metric string, interval time.Duration, upLabels []string) (model.Vector, error) {
	// Metirc has only one data point in the interval and it is present in now time
	// and the up metric has two data points in the interval.
	// This is to ensure that the metric has a missing data point in past not present and the exporter was up in the interval.
	const baseQuery = "count_over_time(%s[%s]) == 1 and %s and on(%s) (count_over_time(up[%s]) == 2)"

	query := fmt.Sprintf(
		baseQuery,
		metric, interval.String(),
		metric,
		strings.Join(upLabels, ","),
		interval.String(),
	)

	// Timeout should be half of the interval
	timeout := interval / 2 //nolint:gomnd,mnd

	result, warnings, err := pq.v1api.Query(ctx, query, time.Now(), v1.WithTimeout(timeout))
	if err != nil {
		return nil, err
	}

	if len(warnings) > 0 {
		slog.WarnContext(ctx, "Prometheus query returned warnings", "query", query, "warnings", warnings)
	}

	v, ok := result.(model.Vector)
	if !ok {
		return nil, fmt.Errorf("unexpected result type: %s", result.Type().String()) //nolint:goerr113
	}

	return v, nil
}
