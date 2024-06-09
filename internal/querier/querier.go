package querier

import (
	"context"
	"time"

	"github.com/prometheus/common/model"
)

type Querier interface {
	Query(ctx context.Context, metric string, interval time.Duration, upLabels []string) (model.Vector, error)
}
